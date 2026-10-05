package repository

import (
	"context"

	"journal-ai-research/backend/internal/models"
)

// ArticleRepository defines the persistence operations for articles and searches
type ArticleRepository interface {
	SaveSearch(ctx context.Context, query string, totalResults int) (string, error)
	SaveArticles(ctx context.Context, articles []models.Article) error
	SaveSearchResults(ctx context.Context, searchID string, articles []models.Article) error
	GetArticleByID(ctx context.Context, id string) (*models.Article, error)
	GetArticlesByIDs(ctx context.Context, ids []string) ([]models.Article, error)
	GetRecentSearches(ctx context.Context, limit int) ([]models.SearchHistory, error)
	Ping(ctx context.Context) error
	Close() error
}
