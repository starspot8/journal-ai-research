package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"journal-ai-research/backend/internal/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(databaseURL string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres connection: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(15 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) Ping(ctx context.Context) error {
	return r.db.PingContext(ctx)
}

func (r *PostgresRepository) Close() error {
	return r.db.Close()
}

func (r *PostgresRepository) SaveSearch(ctx context.Context, query string, totalResults int) (string, error) {
	searchID := uuid.NewString()
	querySQL := `INSERT INTO searches (id, query, total_results, created_at) VALUES ($1, $2, $3, NOW())`
	_, err := r.db.ExecContext(ctx, querySQL, searchID, query, totalResults)
	if err != nil {
		return "", fmt.Errorf("failed to save search: %w", err)
	}
	return searchID, nil
}

func (r *PostgresRepository) SaveArticles(ctx context.Context, articles []models.Article) error {
	if len(articles) == 0 {
		return nil
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO articles (
			id, title, abstract, authors, journal, publication_date, publication_year,
			doi, pmid, pmcid, url, pdf_url, source, is_open_access, full_text, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, NOW()
		) ON CONFLICT (id) DO UPDATE SET
			title = EXCLUDED.title,
			abstract = CASE WHEN LENGTH(EXCLUDED.abstract) > LENGTH(articles.abstract) THEN EXCLUDED.abstract ELSE articles.abstract END,
			authors = CASE WHEN jsonb_array_length(EXCLUDED.authors) > 0 THEN EXCLUDED.authors ELSE articles.authors END,
			journal = COALESCE(NULLIF(EXCLUDED.journal, ''), articles.journal),
			doi = COALESCE(NULLIF(EXCLUDED.doi, ''), articles.doi),
			pmid = COALESCE(NULLIF(EXCLUDED.pmid, ''), articles.pmid),
			pmcid = COALESCE(NULLIF(EXCLUDED.pmcid, ''), articles.pmcid),
			pdf_url = COALESCE(NULLIF(EXCLUDED.pdf_url, ''), articles.pdf_url),
			is_open_access = articles.is_open_access OR EXCLUDED.is_open_access,
			updated_at = NOW()
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare article insert statement: %w", err)
	}
	defer stmt.Close()

	for _, a := range articles {
		authorsJSON, err := json.Marshal(a.Authors)
		if err != nil {
			authorsJSON = []byte("[]")
		}

		_, err = stmt.ExecContext(ctx,
			a.ID,
			a.Title,
			a.Abstract,
			authorsJSON,
			a.Journal,
			a.PublicationDate,
			a.PublicationYear,
			a.DOI,
			a.PMID,
			a.PMCID,
			a.URL,
			a.PDFURL,
			a.Source,
			a.IsOpenAccess,
			a.FullText,
		)
		if err != nil {
			return fmt.Errorf("failed to execute article insert: %w", err)
		}
	}

	return tx.Commit()
}

func (r *PostgresRepository) SaveSearchResults(ctx context.Context, searchID string, articles []models.Article) error {
	if len(articles) == 0 {
		return nil
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO search_results (id, search_id, article_id, rank_position, score, created_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
		ON CONFLICT (search_id, article_id) DO UPDATE SET
			rank_position = EXCLUDED.rank_position,
			score = EXCLUDED.score
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare search result insert: %w", err)
	}
	defer stmt.Close()

	for i, a := range articles {
		srID := uuid.NewString()
		_, err := stmt.ExecContext(ctx, srID, searchID, a.ID, i+1, a.TotalScore)
		if err != nil {
			return fmt.Errorf("failed to insert search result: %w", err)
		}
	}

	return tx.Commit()
}

func (r *PostgresRepository) GetArticleByID(ctx context.Context, id string) (*models.Article, error) {
	query := `
		SELECT id, title, abstract, authors, journal, publication_date, publication_year,
		       doi, pmid, pmcid, url, pdf_url, source, is_open_access, full_text, created_at, updated_at
		FROM articles
		WHERE id = $1 OR doi = $1 OR pmid = $1
		LIMIT 1
	`
	row := r.db.QueryRowContext(ctx, query, id)

	var a models.Article
	var authorsBytes []byte
	var fullText, doi, pmid, pmcid, pdfURL sql.NullString

	err := row.Scan(
		&a.ID,
		&a.Title,
		&a.Abstract,
		&authorsBytes,
		&a.Journal,
		&a.PublicationDate,
		&a.PublicationYear,
		&doi,
		&pmid,
		&pmcid,
		&a.URL,
		&pdfURL,
		&a.Source,
		&a.IsOpenAccess,
		&fullText,
		&a.CreatedAt,
		&a.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to query article: %w", err)
	}

	if doi.Valid {
		a.DOI = doi.String
	}
	if pmid.Valid {
		a.PMID = pmid.String
	}
	if pmcid.Valid {
		a.PMCID = pmcid.String
	}
	if pdfURL.Valid {
		a.PDFURL = pdfURL.String
	}
	if fullText.Valid {
		a.FullText = fullText.String
	}

	if len(authorsBytes) > 0 {
		_ = json.Unmarshal(authorsBytes, &a.Authors)
	}

	return &a, nil
}

func (r *PostgresRepository) GetArticlesByIDs(ctx context.Context, ids []string) ([]models.Article, error) {
	if len(ids) == 0 {
		return []models.Article{}, nil
	}

	query := `
		SELECT id, title, abstract, authors, journal, publication_date, publication_year,
		       doi, pmid, pmcid, url, pdf_url, source, is_open_access, full_text, created_at, updated_at
		FROM articles
		WHERE id = ANY($1)
	`
	rows, err := r.db.QueryContext(ctx, query, pq.Array(ids))
	if err != nil {
		return nil, fmt.Errorf("failed to query articles by ids: %w", err)
	}
	defer rows.Close()

	var list []models.Article
	for rows.Next() {
		var a models.Article
		var authorsBytes []byte
		var fullText, doi, pmid, pmcid, pdfURL sql.NullString

		err := rows.Scan(
			&a.ID,
			&a.Title,
			&a.Abstract,
			&authorsBytes,
			&a.Journal,
			&a.PublicationDate,
			&a.PublicationYear,
			&doi,
			&pmid,
			&pmcid,
			&a.URL,
			&pdfURL,
			&a.Source,
			&a.IsOpenAccess,
			&fullText,
			&a.CreatedAt,
			&a.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan article: %w", err)
		}

		if doi.Valid {
			a.DOI = doi.String
		}
		if pmid.Valid {
			a.PMID = pmid.String
		}
		if pmcid.Valid {
			a.PMCID = pmcid.String
		}
		if pdfURL.Valid {
			a.PDFURL = pdfURL.String
		}
		if fullText.Valid {
			a.FullText = fullText.String
		}
		if len(authorsBytes) > 0 {
			_ = json.Unmarshal(authorsBytes, &a.Authors)
		}

		list = append(list, a)
	}

	return list, nil
}

func (r *PostgresRepository) GetRecentSearches(ctx context.Context, limit int) ([]models.SearchHistory, error) {
	if limit <= 0 {
		limit = 10
	}
	query := `SELECT id, query, total_results, created_at FROM searches ORDER BY created_at DESC LIMIT $1`
	rows, err := r.db.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get recent searches: %w", err)
	}
	defer rows.Close()

	var list []models.SearchHistory
	for rows.Next() {
		var s models.SearchHistory
		if err := rows.Scan(&s.ID, &s.Query, &s.TotalResults, &s.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, nil
}
