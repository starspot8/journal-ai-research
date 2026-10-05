package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"journal-ai-research/backend/internal/models"
	"journal-ai-research/backend/internal/pdf"
	"journal-ai-research/backend/internal/repository"
	"journal-ai-research/backend/internal/search"
)

type Handler struct {
	orchestrator *search.Orchestrator
	repo         repository.ArticleRepository
	startTime    time.Time
}

func NewHandler(orchestrator *search.Orchestrator, repo repository.ArticleRepository) *Handler {
	return &Handler{
		orchestrator: orchestrator,
		repo:         repo,
		startTime:    time.Now(),
	}
}

// HealthHandler returns health status of the service and database
func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	dbStatus := "disabled"
	if h.repo != nil {
		if err := h.repo.Ping(r.Context()); err != nil {
			dbStatus = "unreachable: " + err.Error()
		} else {
			dbStatus = "connected"
		}
	}

	response := map[string]interface{}{
		"status":    "ok",
		"service":   "journal-ai-research-backend",
		"version":   "1.0.0",
		"uptime":    time.Since(h.startTime).String(),
		"database":  dbStatus,
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}

	writeJSON(w, http.StatusOK, response)
}

// SearchHandler processes scientific literature search requests
func (h *Handler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"Method not allowed. Use POST"}`, http.StatusMethodNotAllowed)
		return
	}

	var req models.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON request body"})
		return
	}

	query := strings.TrimSpace(req.Query)
	if len(query) < 2 {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Query too short. Please enter at least 2 characters.",
		})
		return
	}
	if len(query) > 500 {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Query too long. Please limit to 500 characters.",
		})
		return
	}

	limit := req.Limit
	if limit <= 0 || limit > 10 {
		limit = 10
	}

	resp, err := h.orchestrator.Search(r.Context(), query, limit)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "Search failed due to internal error: " + err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, resp)
}

// ArticleDetailHandler retrieves a single article by ID
func (h *Handler) ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, `{"error":"Method not allowed. Use GET"}`, http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/api/v1/articles/")
	id = strings.TrimSuffix(id, "/fulltext")
	id = strings.TrimSpace(id)

	if id == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Article ID required"})
		return
	}

	if h.repo != nil {
		article, err := h.repo.GetArticleByID(r.Context(), id)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Database error: " + err.Error()})
			return
		}
		if article != nil {
			writeJSON(w, http.StatusOK, article)
			return
		}
	}

	writeJSON(w, http.StatusNotFound, map[string]string{"error": "Article not found"})
}

// ArticleFullTextHandler returns legal full-text content or directs to publisher
func (h *Handler) ArticleFullTextHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, `{"error":"Method not allowed. Use GET"}`, http.StatusMethodNotAllowed)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/api/v1/articles/")
	parts := strings.Split(path, "/")
	if len(parts) == 0 || parts[0] == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Article ID required"})
		return
	}
	id := parts[0]

	if h.repo != nil {
		article, err := h.repo.GetArticleByID(r.Context(), id)
		if err == nil && article != nil {
			if !article.IsOpenAccess {
				writeJSON(w, http.StatusForbidden, map[string]interface{}{
					"error":          "Full text is not legally open access through this repository.",
					"is_open_access": false,
					"publisher_url":  article.URL,
					"policy":         "In accordance with project legal policy, restricted articles must be viewed on the publisher site.",
				})
				return
			}

			response := map[string]interface{}{
				"id":             article.ID,
				"title":          article.Title,
				"is_open_access": true,
				"full_text":      article.FullText,
				"pdf_url":        article.PDFURL,
				"original_url":   article.URL,
			}
			writeJSON(w, http.StatusOK, response)
			return
		}
	}

	writeJSON(w, http.StatusNotFound, map[string]string{"error": "Article not found"})
}

// ExportPDFHandler generates a PDF research report for selected articles
func (h *Handler) ExportPDFHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"Method not allowed. Use POST"}`, http.StatusMethodNotAllowed)
		return
	}

	var req models.ExportPDFRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON request body"})
		return
	}

	if len(req.ArticleIDs) == 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "At least one article ID required"})
		return
	}

	var articles []models.Article
	if h.repo != nil {
		var err error
		articles, err = h.repo.GetArticlesByIDs(r.Context(), req.ArticleIDs)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to fetch articles: " + err.Error()})
			return
		}
	}

	if len(articles) == 0 {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "No matching articles found in repository"})
		return
	}

	pdfBytes, err := pdf.GenerateResearchReport(req.Query, articles)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to generate PDF: " + err.Error()})
		return
	}

	filename := fmt.Sprintf("journal-ai-research-report-%d.pdf", time.Now().Unix())
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", filename))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(pdfBytes)))
	w.WriteHeader(http.StatusOK)
	w.Write(pdfBytes)
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
