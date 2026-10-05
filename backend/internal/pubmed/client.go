package pubmed

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
	apiKey     string
}

func NewClient(httpClient *http.Client, apiKey string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 8 * time.Second}
	}
	return &Client{
		httpClient: httpClient,
		apiKey:     apiKey,
	}
}

type esearchResponse struct {
	ESearchResult struct {
		IDList []string `json:"idlist"`
		Count  string   `json:"count"`
	} `json:"esearchresult"`
}

type esummaryResponse struct {
	Result map[string]interface{} `json:"result"`
}

func (c *Client) Search(ctx context.Context, query string, limit int) ([]models.Article, error) {
	if limit <= 0 {
		limit = 10
	}

	searchURL := fmt.Sprintf("https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esearch.fcgi?db=pubmed&term=%s&retmode=json&retmax=%d&sort=pub_date",
		url.QueryEscape(query), limit)
	if c.apiKey != "" {
		searchURL += "&api_key=" + url.QueryEscape(c.apiKey)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, searchURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create pubmed search request: %w", err)
	}
	req.Header.Set("User-Agent", "JournalAIResearch/1.0 (academic search)")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("pubmed search request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("pubmed search returned status %d", resp.StatusCode)
	}

	var searchData esearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchData); err != nil {
		return nil, fmt.Errorf("failed to decode pubmed search response: %w", err)
	}

	idList := searchData.ESearchResult.IDList
	if len(idList) == 0 {
		return []models.Article{}, nil
	}

	// Fetch summaries for retrieved IDs
	summaryURL := fmt.Sprintf("https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esummary.fcgi?db=pubmed&id=%s&retmode=json",
		strings.Join(idList, ","))
	if c.apiKey != "" {
		summaryURL += "&api_key=" + url.QueryEscape(c.apiKey)
	}

	summaryReq, err := http.NewRequestWithContext(ctx, http.MethodGet, summaryURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create pubmed summary request: %w", err)
	}
	summaryReq.Header.Set("User-Agent", "JournalAIResearch/1.0 (academic search)")

	summaryResp, err := c.httpClient.Do(summaryReq)
	if err != nil {
		return nil, fmt.Errorf("pubmed summary request failed: %w", err)
	}
	defer summaryResp.Body.Close()

	if summaryResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("pubmed summary returned status %d", summaryResp.StatusCode)
	}

	var summaryData esummaryResponse
	if err := json.NewDecoder(summaryResp.Body).Decode(&summaryData); err != nil {
		return nil, fmt.Errorf("failed to decode pubmed summary response: %w", err)
	}

	var articles []models.Article
	for _, pmid := range idList {
		itemRaw, ok := summaryData.Result[pmid]
		if !ok {
			continue
		}
		itemMap, ok := itemRaw.(map[string]interface{})
		if !ok {
			continue
		}

		title, _ := itemMap["title"].(string)
		title = strings.TrimSpace(strings.TrimSuffix(title, "."))
		sourceJournal, _ := itemMap["source"].(string)
		pubDate, _ := itemMap["pubdate"].(string)
		
		var pubYear int
		if len(pubDate) >= 4 {
			if y, err := strconv.Atoi(pubDate[:4]); err == nil {
				pubYear = y
			}
		}

		// Extract authors
		var authors []models.Author
		if authorsRaw, ok := itemMap["authors"].([]interface{}); ok {
			for _, a := range authorsRaw {
				if aMap, ok := a.(map[string]interface{}); ok {
					if name, ok := aMap["name"].(string); ok && name != "" {
						authors = append(authors, models.Author{Name: name})
					}
				}
			}
		}

		// Extract article IDs (DOI, PMCID)
		var doi, pmcid string
		if articleIDsRaw, ok := itemMap["articleids"].([]interface{}); ok {
			for _, idObj := range articleIDsRaw {
				if idMap, ok := idObj.(map[string]interface{}); ok {
					idType, _ := idMap["idtype"].(string)
					val, _ := idMap["value"].(string)
					if idType == "doi" {
						doi = val
					} else if idType == "pmc" || idType == "pmcid" {
						pmcid = val
					}
				}
			}
		}

		articleURL := fmt.Sprintf("https://pubmed.ncbi.nlm.nih.gov/%s/", pmid)
		if doi != "" {
			articleURL = fmt.Sprintf("https://doi.org/%s", doi)
		}

		article := models.Article{
			ID:              "pmid-" + pmid,
			Title:           title,
			Abstract:        fmt.Sprintf("Indexed in PubMed (PMID: %s). Published in %s.", pmid, sourceJournal),
			Authors:         authors,
			Journal:         sourceJournal,
			PublicationDate: pubDate,
			PublicationYear: pubYear,
			DOI:             doi,
			PMID:            pmid,
			PMCID:           pmcid,
			URL:             articleURL,
			Source:          "pubmed",
			IsOpenAccess:    pmcid != "",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}
		if pmcid != "" {
			article.PDFURL = fmt.Sprintf("https://www.ncbi.nlm.nih.gov/pmc/articles/%s/pdf/", pmcid)
		}

		articles = append(articles, article)
	}

	return articles, nil
}
