package models

import "time"

// Author represents an author of a scientific publication
type Author struct {
	Name        string `json:"name"`
	Affiliation string `json:"affiliation,omitempty"`
}

// Article represents a normalized scientific journal article
type Article struct {
	ID                string    `json:"id"`
	Title             string    `json:"title"`
	Abstract          string    `json:"abstract"`
	Authors           []Author  `json:"authors"`
	Journal           string    `json:"journal"`
	PublicationDate   string    `json:"publication_date"`
	PublicationYear   int       `json:"publication_year"`
	DOI               string    `json:"doi,omitempty"`
	PMID              string    `json:"pmid,omitempty"`
	PMCID             string    `json:"pmcid,omitempty"`
	URL               string    `json:"url"`
	PDFURL            string    `json:"pdf_url,omitempty"`
	Source            string    `json:"source"` // "pubmed", "europepmc", "crossref", "composite"
	IsOpenAccess      bool      `json:"is_open_access"`
	FullText          string    `json:"full_text,omitempty"`
	RelevanceScore    float64   `json:"relevance_score"`
	RecencyScore      float64   `json:"recency_score"`
	CompletenessScore float64   `json:"completeness_score"`
	TotalScore        float64   `json:"total_score"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// SearchRequest represents the payload for POST /api/v1/search
type SearchRequest struct {
	Query string `json:"query"`
	Limit int    `json:"limit,omitempty"`
}

// SearchResponse represents the response for POST /api/v1/search
type SearchResponse struct {
	Query     string         `json:"query"`
	Results   []Article      `json:"results"`
	Total     int            `json:"total"`
	SearchID  string         `json:"search_id"`
	Sources   map[string]int `json:"sources"`
	Warnings  []string       `json:"warnings,omitempty"`
	Timestamp string         `json:"timestamp"`
}

// ExportPDFRequest represents the payload for POST /api/v1/export/pdf
type ExportPDFRequest struct {
	Query      string   `json:"query"`
	ArticleIDs []string `json:"article_ids"`
}

// SearchHistory represents an entry in the search audit table
type SearchHistory struct {
	ID           string    `json:"id"`
	Query        string    `json:"query"`
	TotalResults int       `json:"total_results"`
	CreatedAt    time.Time `json:"created_at"`
}
