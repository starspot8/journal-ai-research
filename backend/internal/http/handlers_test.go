package http

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"journal-ai-research/backend/internal/models"
	"journal-ai-research/backend/internal/search"
)

type mockRepository struct{}

func (m *mockRepository) SaveSearch(ctx context.Context, query string, totalResults int) (string, error) {
	return "mock-search-id", nil
}
func (m *mockRepository) SaveArticles(ctx context.Context, articles []models.Article) error {
	return nil
}
func (m *mockRepository) SaveSearchResults(ctx context.Context, searchID string, articles []models.Article) error {
	return nil
}
func (m *mockRepository) GetArticleByID(ctx context.Context, id string) (*models.Article, error) {
	if id == "art-test-1" {
		return &models.Article{
			ID:           "art-test-1",
			Title:        "Test Article",
			IsOpenAccess: true,
			FullText:     "Complete open access research contents.",
		}, nil
	}
	return nil, nil
}
func (m *mockRepository) GetArticlesByIDs(ctx context.Context, ids []string) ([]models.Article, error) {
	var list []models.Article
	for _, id := range ids {
		if id == "art-test-1" {
			list = append(list, models.Article{ID: id, Title: "Test Article"})
		}
	}
	return list, nil
}
func (m *mockRepository) GetRecentSearches(ctx context.Context, limit int) ([]models.SearchHistory, error) {
	return []models.SearchHistory{}, nil
}
func (m *mockRepository) Ping(ctx context.Context) error {
	return nil
}
func (m *mockRepository) Close() error {
	return nil
}

func TestHealthHandler(t *testing.T) {
	handler := NewHandler(nil, nil)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	w := httptest.NewRecorder()

	handler.HealthHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	body := w.Body.String()
	if !bytes.Contains([]byte(body), []byte(`"status":"ok"`)) {
		t.Errorf("expected response to contain status ok, got %s", body)
	}
}

func TestSearchHandler_Validation(t *testing.T) {
	orchestrator := search.NewOrchestrator(nil, nil, nil, nil, 0)
	handler := NewHandler(orchestrator, nil)

	// Too short query
	req := httptest.NewRequest(http.MethodPost, "/api/v1/search", bytes.NewBufferString(`{"query":"a"}`))
	w := httptest.NewRecorder()

	handler.SearchHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for short query, got %d", w.Code)
	}
}
