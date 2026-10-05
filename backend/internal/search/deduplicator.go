package search

import (
	"strings"

	"journal-ai-research/backend/internal/models"
)

// DeduplicateArticles removes duplicates across scientific sources and merges complementary fields
func DeduplicateArticles(articles []models.Article) []models.Article {
	doiIndex := make(map[string]int)
	pmidIndex := make(map[string]int)
	pmcidIndex := make(map[string]int)
	titleIndex := make(map[string]int)

	var unique []models.Article

	for _, a := range articles {
		a = NormalizeArticle(a)
		if a.Title == "" {
			continue
		}

		canonTitle := CanonicalTitle(a.Title)
		existingIdx := -1

		if a.DOI != "" {
			if idx, ok := doiIndex[a.DOI]; ok {
				existingIdx = idx
			}
		}
		if existingIdx == -1 && a.PMID != "" {
			if idx, ok := pmidIndex[a.PMID]; ok {
				existingIdx = idx
			}
		}
		if existingIdx == -1 && a.PMCID != "" {
			if idx, ok := pmcidIndex[a.PMCID]; ok {
				existingIdx = idx
			}
		}
		if existingIdx == -1 && len(canonTitle) >= 20 {
			if idx, ok := titleIndex[canonTitle]; ok {
				existingIdx = idx
			}
		}

		if existingIdx != -1 {
			// Merge into existing article
			merged := mergeArticles(unique[existingIdx], a)
			unique[existingIdx] = merged
		} else {
			// Insert as new article
			newIdx := len(unique)
			unique = append(unique, a)

			if a.DOI != "" {
				doiIndex[a.DOI] = newIdx
			}
			if a.PMID != "" {
				pmidIndex[a.PMID] = newIdx
			}
			if a.PMCID != "" {
				pmcidIndex[a.PMCID] = newIdx
			}
			if len(canonTitle) >= 20 {
				titleIndex[canonTitle] = newIdx
			}
		}
	}

	return unique
}

func mergeArticles(existing, incoming models.Article) models.Article {
	// Source tracking
	if existing.Source != incoming.Source && !strings.Contains(existing.Source, "composite") {
		existing.Source = "composite (" + existing.Source + "+" + incoming.Source + ")"
	}

	// Prefer longer abstract
	if len(incoming.Abstract) > len(existing.Abstract) {
		existing.Abstract = incoming.Abstract
	}

	// Merge identifiers
	if existing.DOI == "" && incoming.DOI != "" {
		existing.DOI = incoming.DOI
	}
	if existing.PMID == "" && incoming.PMID != "" {
		existing.PMID = incoming.PMID
	}
	if existing.PMCID == "" && incoming.PMCID != "" {
		existing.PMCID = incoming.PMCID
	}

	// Prefer more complete metadata
	if len(existing.Authors) == 0 && len(incoming.Authors) > 0 {
		existing.Authors = incoming.Authors
	}
	if existing.Journal == "" && incoming.Journal != "" {
		existing.Journal = incoming.Journal
	}
	if existing.PublicationYear == 0 && incoming.PublicationYear != 0 {
		existing.PublicationYear = incoming.PublicationYear
		existing.PublicationDate = incoming.PublicationDate
	}

	// Open Access & PDF
	if incoming.IsOpenAccess {
		existing.IsOpenAccess = true
	}
	if existing.PDFURL == "" && incoming.PDFURL != "" {
		existing.PDFURL = incoming.PDFURL
	}
	if existing.FullText == "" && incoming.FullText != "" {
		existing.FullText = incoming.FullText
	}

	return existing
}
