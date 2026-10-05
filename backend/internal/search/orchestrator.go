package search

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"journal-ai-research/backend/internal/crossref"
	"journal-ai-research/backend/internal/europepmc"
	"journal-ai-research/backend/internal/models"
	"journal-ai-research/backend/internal/pubmed"
	"journal-ai-research/backend/internal/ranking"
)

type Repository interface {
	SaveSearch(ctx context.Context, query string, totalResults int) (string, error)
	SaveArticles(ctx context.Context, articles []models.Article) error
	SaveSearchResults(ctx context.Context, searchID string, articles []models.Article) error
}

type Orchestrator struct {
	pubmedClient    *pubmed.Client
	europepmcClient *europepmc.Client
	crossrefClient  *crossref.Client
	repo            Repository
	searchTimeout   time.Duration
}

func NewOrchestrator(
	p *pubmed.Client,
	e *europepmc.Client,
	c *crossref.Client,
	repo Repository,
	searchTimeout time.Duration,
) *Orchestrator {
	if searchTimeout <= 0 {
		searchTimeout = 10 * time.Second
	}
	return &Orchestrator{
		pubmedClient:    p,
		europepmcClient: e,
		crossrefClient:  c,
		repo:            repo,
		searchTimeout:   searchTimeout,
	}
}

func (o *Orchestrator) Search(ctx context.Context, query string, limit int) (*models.SearchResponse, error) {
	if limit <= 0 {
		limit = 10
	}

	searchCtx, cancel := context.WithTimeout(ctx, o.searchTimeout)
	defer cancel()

	var wg sync.WaitGroup
	var mu sync.Mutex

	var allArticles []models.Article
	sourceCounts := make(map[string]int)
	var warnings []string

	// 1. PubMed
	wg.Add(1)
	go func() {
		defer wg.Done()
		articles, err := o.pubmedClient.Search(searchCtx, query, limit)
		mu.Lock()
		defer mu.Unlock()
		if err != nil {
			warnings = append(warnings, fmt.Sprintf("PubMed provider error: %v", err))
		} else {
			sourceCounts["pubmed"] = len(articles)
			allArticles = append(allArticles, articles...)
		}
	}()

	// 2. Europe PMC
	wg.Add(1)
	go func() {
		defer wg.Done()
		articles, err := o.europepmcClient.Search(searchCtx, query, limit)
		mu.Lock()
		defer mu.Unlock()
		if err != nil {
			warnings = append(warnings, fmt.Sprintf("Europe PMC provider error: %v", err))
		} else {
			sourceCounts["europepmc"] = len(articles)
			allArticles = append(allArticles, articles...)
		}
	}()

	// 3. Crossref
	wg.Add(1)
	go func() {
		defer wg.Done()
		articles, err := o.crossrefClient.Search(searchCtx, query, limit)
		mu.Lock()
		defer mu.Unlock()
		if err != nil {
			warnings = append(warnings, fmt.Sprintf("Crossref provider error: %v", err))
		} else {
			sourceCounts["crossref"] = len(articles)
			allArticles = append(allArticles, articles...)
		}
	}()

	wg.Wait()

	// Normalization & Deduplication
	deduped := DeduplicateArticles(allArticles)

	// Ranking
	ranked := ranking.RankArticles(query, deduped, limit)

	searchID := uuid.NewString()

	// Persist to database if repository is available
	if o.repo != nil {
		saveCtx, saveCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer saveCancel()

		if sID, err := o.repo.SaveSearch(saveCtx, query, len(ranked)); err == nil && sID != "" {
			searchID = sID
		}
		if len(ranked) > 0 {
			_ = o.repo.SaveArticles(saveCtx, ranked)
			_ = o.repo.SaveSearchResults(saveCtx, searchID, ranked)
		}
	}

	response := &models.SearchResponse{
		Query:     query,
		Results:   ranked,
		Total:     len(ranked),
		SearchID:  searchID,
		Sources:   sourceCounts,
		Warnings:  warnings,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	return response, nil
}
