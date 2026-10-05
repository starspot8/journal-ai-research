package http

import (
	"net/http"
	"strings"

	"journal-ai-research/backend/internal/config"
)

// NewRouter registers application routes and wraps them with required middlewares
func NewRouter(cfg *config.Config, handler *Handler) http.Handler {
	mux := http.NewServeMux()

	// Healthcheck
	mux.HandleFunc("/api/v1/health", handler.HealthHandler)

	// Search
	mux.HandleFunc("/api/v1/search", handler.SearchHandler)

	// Article detail & Fulltext
	mux.HandleFunc("/api/v1/articles/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/fulltext") {
			handler.ArticleFullTextHandler(w, r)
		} else {
			handler.ArticleDetailHandler(w, r)
		}
	})

	// PDF Export
	mux.HandleFunc("/api/v1/export/pdf", handler.ExportPDFHandler)

	// Rate limiter
	rateLimiter := NewRateLimiter(cfg.RateLimitPerMinute)

	// Compose middleware chain
	var rootHandler http.Handler = mux
	rootHandler = rateLimiter.Middleware(rootHandler)
	rootHandler = SecurityHeaders(rootHandler)
	rootHandler = CORSMiddleware(cfg.AllowedOrigins)(rootHandler)
	rootHandler = RequestLogger(rootHandler)

	return rootHandler
}
