package crossref

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"journal-ai-research/backend/internal/models"
)

type Client struct {
	httpClient *http.Client
	mailto     string
}

func NewClient(httpClient *http.Client, mailto string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 8 * time.Second}
	}
	return &Client{
		httpClient: httpClient,
		mailto:     mailto,
	}
}

type crossrefResponse struct {
	Message struct {
		TotalResults int `json:"total-results"`
		Items        []struct {
			DOI            string   `json:"DOI"`
			Title          []string `json:"title"`
			ContainerTitle []string `json:"container-title"`
			Author         []struct {
				Given  string `json:"given"`
				Family string `json:"family"`
				Name   string `json:"name"`
			} `json:"author"`
			Published struct {
				DateParts [][]int `json:"date-parts"`
			} `json:"published"`
			URL  string `json:"URL"`
			Link []struct {
				URL                 string `json:"URL"`
				ContentType         string `json:"content-type"`
				ContentVersion      string `json:"content-version"`
				IntendedApplication string `json:"intended-application"`
			} `json:"link"`
			Abstract string `json:"abstract"`
		} `json:"items"`
	} `json:"message"`
}

func (c *Client) Search(ctx context.Context, query string, limit int) ([]models.Article, error) {
	if limit <= 0 {
		limit = 10
	}

	searchURL := fmt.Sprintf("https://api.crossref.org/works?query=%s&rows=%d&sort=published&order=desc",
		url.QueryEscape(query), limit)
	if c.mailto != "" {
		searchURL += "&mailto=" + url.QueryEscape(c.mailto)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, searchURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create crossref request: %w", err)
	}
	req.Header.Set("User-Agent", fmt.Sprintf("JournalAIResearch/1.0 (mailto:%s)", c.mailto))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("crossref request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("crossref returned status %d", resp.StatusCode)
	}

	var data crossrefResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode crossref response: %w", err)
	}

	var articles []models.Article
	for _, item := range data.Message.Items {
		var title string
		if len(item.Title) > 0 {
			title = strings.TrimSpace(item.Title[0])
		}
		if title == "" {
			continue
		}

		var journal string
		if len(item.ContainerTitle) > 0 {
			journal = strings.TrimSpace(item.ContainerTitle[0])
		}

		var pubYear int
		var pubDate string
		if len(item.Published.DateParts) > 0 && len(item.Published.DateParts[0]) > 0 {
			pubYear = item.Published.DateParts[0][0]
			pubDate = fmt.Sprintf("%d", pubYear)
			if len(item.Published.DateParts[0]) > 1 {
				pubDate = fmt.Sprintf("%d-%02d", pubYear, item.Published.DateParts[0][1])
			}
		}

		var authors []models.Author
		for _, a := range item.Author {
			name := strings.TrimSpace(a.Name)
			if name == "" {
				name = strings.TrimSpace(fmt.Sprintf("%s %s", a.Given, a.Family))
			}
			if name != "" {
				authors = append(authors, models.Author{Name: name})
			}
		}

		var pdfURL string
		isOpenAccess := false
		for _, l := range item.Link {
			if strings.Contains(l.ContentType, "pdf") {
				pdfURL = l.URL
				isOpenAccess = true
				break
			}
		}

		articleURL := item.URL
		if articleURL == "" && item.DOI != "" {
			articleURL = fmt.Sprintf("https://doi.org/%s", item.DOI)
		}

		abstract := cleanAbstract(item.Abstract)
		if abstract == "" {
			abstract = fmt.Sprintf("Published in %s (%d). DOI: %s.", journal, pubYear, item.DOI)
		}

		articleID := "crossref-" + strings.ReplaceAll(item.DOI, "/", "_")
		if item.DOI == "" {
			articleID = fmt.Sprintf("crossref-%d", time.Now().UnixNano())
		}

		articles = append(articles, models.Article{
			ID:              articleID,
			Title:           title,
			Abstract:        abstract,
			Authors:         authors,
			Journal:         journal,
			PublicationDate: pubDate,
			PublicationYear: pubYear,
			DOI:             item.DOI,
			URL:             articleURL,
			PDFURL:          pdfURL,
			Source:          "crossref",
			IsOpenAccess:    isOpenAccess,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		})
	}

	return articles, nil
}

func cleanAbstract(text string) string {
	text = strings.ReplaceAll(text, "<jats:p>", "")
	text = strings.ReplaceAll(text, "</jats:p>", "\n")
	text = strings.ReplaceAll(text, "<jats:title>", "")
	text = strings.ReplaceAll(text, "</jats:title>", "\n")
	text = strings.ReplaceAll(text, "<p>", "")
	text = strings.ReplaceAll(text, "</p>", "\n")
	return strings.TrimSpace(text)
}
