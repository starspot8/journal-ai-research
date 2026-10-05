package europepmc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"journal-ai-research/backend/internal/models"
)

type Client struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 8 * time.Second}
	}
	return &Client{httpClient: httpClient}
}

type ResultItem struct {
	ID           string `json:"id"`
	PMID         string `json:"pmid"`
	PMCID        string `json:"pmcid"`
	DOI          string `json:"doi"`
	Title        string `json:"title"`
	AuthorString string `json:"authorString"`
	JournalTitle string `json:"journalTitle"`
	PubYear      string `json:"pubYear"`
	AbstractText string `json:"abstractText"`
	IsOpenAccess string `json:"isOpenAccess"`
}

type europePMCResponse struct {
	HitCount   int `json:"hitCount"`
	ResultList struct {
		Result []ResultItem `json:"result"`
	} `json:"resultList"`
}

func (c *Client) Search(ctx context.Context, query string, limit int) ([]models.Article, error) {
	if limit <= 0 {
		limit = 10
	}

	searchURL := fmt.Sprintf("https://www.ebi.ac.uk/europepmc/webservices/rest/search?query=%s&format=json&pageSize=%d&resultType=core&sort=P_PD_D%%20desc",
		url.QueryEscape(query), limit)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, searchURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create europepmc request: %w", err)
	}
	req.Header.Set("User-Agent", "JournalAIResearch/1.0 (academic search)")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("europepmc request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("europepmc returned status %d", resp.StatusCode)
	}

	var data europePMCResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode europepmc response: %w", err)
	}

	var articles []models.Article
	for _, item := range data.ResultList.Result {
		title := strings.TrimSpace(strings.TrimSuffix(item.Title, "."))
		if title == "" {
			continue
		}

		var pubYear int
		if item.PubYear != "" {
			pubYear, _ = strconv.Atoi(item.PubYear)
		}

		var authors []models.Author
		if item.AuthorString != "" {
			parts := strings.Split(item.AuthorString, ",")
			for _, p := range parts {
				p = strings.TrimSpace(p)
				if p != "" {
					authors = append(authors, models.Author{Name: p})
				}
			}
		}

		isOpenAccess := strings.EqualFold(item.IsOpenAccess, "Y")
		articleID := item.ID
		if articleID == "" {
			if item.PMID != "" {
				articleID = "pmid-" + item.PMID
			} else if item.DOI != "" {
				articleID = "doi-" + strings.ReplaceAll(item.DOI, "/", "_")
			} else {
				articleID = fmt.Sprintf("epmc-%d", time.Now().UnixNano())
			}
		} else {
			articleID = "epmc-" + articleID
		}

		articleURL := fmt.Sprintf("https://europepmc.org/article/%s/%s", item.SourceOrDefault(), item.ID)
		if item.DOI != "" {
			articleURL = fmt.Sprintf("https://doi.org/%s", item.DOI)
		}

		var pdfURL string
		if item.PMCID != "" {
			pdfURL = fmt.Sprintf("https://europepmc.org/backend/ptpmcrender.fcgi?accid=%s&blobtype=pdf", item.PMCID)
		}

		abstract := cleanAbstract(item.AbstractText)
		if abstract == "" {
			abstract = fmt.Sprintf("Article published in %s (%d). Europe PMC Record ID: %s.", item.JournalTitle, pubYear, item.ID)
		}

		articles = append(articles, models.Article{
			ID:              articleID,
			Title:           title,
			Abstract:        abstract,
			Authors:         authors,
			Journal:         item.JournalTitle,
			PublicationDate: item.PubYear,
			PublicationYear: pubYear,
			DOI:             item.DOI,
			PMID:            item.PMID,
			PMCID:           item.PMCID,
			URL:             articleURL,
			PDFURL:          pdfURL,
			Source:          "europepmc",
			IsOpenAccess:    isOpenAccess,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		})
	}

	return articles, nil
}

func (r europePMCResponse) HitCountOrDefault() int {
	return r.HitCount
}

func cleanAbstract(text string) string {
	text = strings.ReplaceAll(text, "<p>", "")
	text = strings.ReplaceAll(text, "</p>", "\n")
	text = strings.ReplaceAll(text, "<b>", "")
	text = strings.ReplaceAll(text, "</b>", "")
	text = strings.ReplaceAll(text, "<i>", "")
	text = strings.ReplaceAll(text, "</i>", "")
	return strings.TrimSpace(text)
}

func (r ResultItem) SourceOrDefault() string {
	if r.PMCID != "" {
		return "PMC"
	}
	return "MED"
}
