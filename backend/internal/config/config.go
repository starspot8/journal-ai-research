package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Port               string
	Environment        string
	DatabaseURL        string
	HTTPTimeout        time.Duration
	SearchTimeout      time.Duration
	RateLimitPerMinute int
	AllowedOrigins     []string
	NCBIAPIKey         string
	CrossrefMailto     string
}

func LoadConfig() *Config {
	port := getEnv("PORT", "8080")
	env := getEnv("ENVIRONMENT", "development")
	dbURL := getEnv("DATABASE_URL", "postgres://journal_user:journal_password@127.0.0.1:5435/journal_ai_db?sslmode=disable")
	
	httpTimeoutSec := getEnvAsInt("HTTP_TIMEOUT_SECONDS", 15)
	searchTimeoutSec := getEnvAsInt("SEARCH_TIMEOUT_SECONDS", 10)
	rateLimit := getEnvAsInt("RATE_LIMIT_PER_MINUTE", 60)

	originsRaw := getEnv("ALLOWED_ORIGINS", "http://localhost:3000,http://127.0.0.1:3000,http://localhost:5173")
	var allowedOrigins []string
	for _, o := range strings.Split(originsRaw, ",") {
		o = strings.TrimSpace(o)
		if o != "" {
			allowedOrigins = append(allowedOrigins, o)
		}
	}

	return &Config{
		Port:               port,
		Environment:        env,
		DatabaseURL:        dbURL,
		HTTPTimeout:        time.Duration(httpTimeoutSec) * time.Second,
		SearchTimeout:      time.Duration(searchTimeoutSec) * time.Second,
		RateLimitPerMinute: rateLimit,
		AllowedOrigins:     allowedOrigins,
		NCBIAPIKey:         os.Getenv("NCBI_API_KEY"),
		CrossrefMailto:     getEnv("CROSSREF_MAILTO", "research@journal-ai.local"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return fallback
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return fallback
	}
	return val
}
