package search

import (
	"regexp"
	"strings"

	"journal-ai-research/backend/internal/models"
)

var (
	htmlTagRegex = regexp.MustCompile(`<[^>]*>`)
	spaceRegex   = regexp.MustCompile(`\s+`)
)

// NormalizeArticle cleans and standardizes metadata fields in an Article
func NormalizeArticle(a models.Article) models.Article {
	a.Title = CleanString(a.Title)
	a.Abstract = CleanString(a.Abstract)
	a.Journal = CleanString(a.Journal)
	a.DOI = strings.TrimSpace(strings.ToLower(a.DOI))
	a.PMID = strings.TrimSpace(a.PMID)
	a.PMCID = strings.TrimSpace(a.PMCID)

	// Clean authors
	for i := range a.Authors {
		a.Authors[i].Name = CleanString(a.Authors[i].Name)
	}

	// Clean URL
	a.URL = strings.TrimSpace(a.URL)
	a.PDFURL = strings.TrimSpace(a.PDFURL)

	return a
}

// CleanString strips HTML tags and normalizes consecutive spaces
func CleanString(s string) string {
	s = htmlTagRegex.ReplaceAllString(s, " ")
	s = spaceRegex.ReplaceAllString(s, " ")
	return strings.TrimSpace(s)
}

// CanonicalTitle creates a lowercase alphanumeric representation for fuzzy duplicate matching
func CanonicalTitle(title string) string {
	title = strings.ToLower(title)
	var sb strings.Builder
	for _, r := range title {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
