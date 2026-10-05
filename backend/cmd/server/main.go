package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"journal-ai-research/backend/internal/config"
	"journal-ai-research/backend/internal/crossref"
	"journal-ai-research/backend/internal/europepmc"
	httpLayer "journal-ai-research/backend/internal/http"
	"journal-ai-research/backend/internal/pubmed"
	"journal-ai-research/backend/internal/repository"
	"journal-ai-research/backend/internal/search"
)

func main() {
	cfg := config.LoadConfig()

	log.Printf("[Journal AI Research] Starting backend service on port %s (env: %s)...", cfg.Port, cfg.Environment)

	// Initialize PostgreSQL repository
	var repo repository.ArticleRepository
	pgRepo, err := repository.NewPostgresRepository(cfg.DatabaseURL)
	if err != nil {
		log.Printf("[WARN] PostgreSQL connection failed (%v). Running without database persistence.", err)
	} else {
		repo = pgRepo
		defer pgRepo.Close()
		log.Printf("[DB] Connected to PostgreSQL successfully at %s", cfg.DatabaseURL)
	}

	// Shared HTTP client for external providers with configured timeout
	client := &http.Client{
		Timeout: cfg.HTTPTimeout,
	}

	// Initialize scientific providers
	pubmedClient := pubmed.NewClient(client, cfg.NCBIAPIKey)
	europepmcClient := europepmc.NewClient(client)
	crossrefClient := crossref.NewClient(client, cfg.CrossrefMailto)

	// Initialize search orchestrator
	orchestrator := search.NewOrchestrator(
		pubmedClient,
		europepmcClient,
		crossrefClient,
		repo,
		cfg.SearchTimeout,
	)

	// Initialize HTTP handler and router
	handler := httpLayer.NewHandler(orchestrator, repo)
	router := httpLayer.NewRouter(cfg, handler)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      router,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Channel for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("[HTTP] Server listening on http://localhost:%s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[FATAL] HTTP server error: %v", err)
		}
	}()

	<-stop
	log.Println("[SHUTDOWN] Received signal to stop. Shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("[ERROR] Server forced to shutdown: %v", err)
	}

	log.Println("[SHUTDOWN] Journal AI Research backend exited cleanly.")
}
