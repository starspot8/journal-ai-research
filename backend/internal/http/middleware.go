package http

import (
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// RequestLogger logs incoming HTTP requests with response status and duration
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(wrapper, r)

		log.Printf("[HTTP] %s %s | Status: %d | Duration: %v | IP: %s",
			r.Method, r.URL.Path, wrapper.statusCode, time.Since(start), r.RemoteAddr)
	})
}

// SecurityHeaders applies standard web security headers
func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		next.ServeHTTP(w, r)
	})
}

// CORSMiddleware enables cross-origin resource sharing for configured origins
func CORSMiddleware(allowedOrigins []string) func(http.Handler) http.Handler {
	originsMap := make(map[string]bool)
	for _, o := range allowedOrigins {
		originsMap[strings.TrimSpace(o)] = true
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if origin != "" && (originsMap[origin] || originsMap["*"] || len(originsMap) == 0) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
				w.Header().Set("Access-Control-Max-Age", "86400")
			}

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// RateLimiter limits requests per client IP within a rolling 1-minute window
type RateLimiter struct {
	mu      sync.Mutex
	clients map[string]*clientBucket
	limit   int
}

type clientBucket struct {
	count     int
	resetTime time.Time
}

func NewRateLimiter(limitPerMinute int) *RateLimiter {
	if limitPerMinute <= 0 {
		limitPerMinute = 60
	}
	rl := &RateLimiter{
		clients: make(map[string]*clientBucket),
		limit:   limitPerMinute,
	}

	// Periodic cleanup of expired buckets
	go func() {
		ticker := time.NewTicker(2 * time.Minute)
		for range ticker.C {
			rl.mu.Lock()
			now := time.Now()
			for ip, b := range rl.clients {
				if now.After(b.resetTime) {
					delete(rl.clients, ip)
				}
			}
			rl.mu.Unlock()
		}
	}()

	return rl
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]

		rl.mu.Lock()
		bucket, exists := rl.clients[ip]
		now := time.Now()

		if !exists || now.After(bucket.resetTime) {
			rl.clients[ip] = &clientBucket{
				count:     1,
				resetTime: now.Add(time.Minute),
			}
			rl.mu.Unlock()
			next.ServeHTTP(w, r)
			return
		}

		if bucket.count >= rl.limit {
			rl.mu.Unlock()
			http.Error(w, `{"error":"Too many requests. Please wait a moment."}`, http.StatusTooManyRequests)
			return
		}

		bucket.count++
		rl.mu.Unlock()

		next.ServeHTTP(w, r)
	})
}
