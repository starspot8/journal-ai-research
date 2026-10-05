package ranking

import (
	"sort"
	"strings"
	"time"

	"journal-ai-research/backend/internal/models"
)

var stopWords = map[string]bool{
	"the": true, "and": true, "of": true, "to": true, "in": true, "a": true, "is": true, "that": true, "for": true, "on": true, "with": true,
	"as": true, "by": true, "at": true, "from": true, "an": true, "be": true, "this": true, "which": true, "or": true,
	// Indonesian common stop words
	"dan": true, "di": true, "ke": true, "dari": true, "yang": true, "pada": true, "untuk": true, "dengan": true, "adalah": true, "ini": true, "itu": true, "dalam": true,
}

// RankArticles calculates relevance (60%), recency (30%), and metadata completeness (10%),
// returning the top N ranked articles.
func RankArticles(query string, articles []models.Article, limit int) []models.Article {
	if limit <= 0 {
		limit = 10
	}

	queryTokens := tokenize(query)
	currentYear := time.Now().Year()

	ranked := make([]models.Article, len(articles))
	copy(ranked, articles)

	for i := range ranked {
		relevance := calculateRelevance(query, queryTokens, &ranked[i])
		recency := calculateRecency(ranked[i].PublicationYear, currentYear)
		completeness := calculateCompleteness(&ranked[i])

		total := (0.60 * relevance) + (0.30 * recency) + (0.10 * completeness)

		ranked[i].RelevanceScore = round4(relevance)
		ranked[i].RecencyScore = round4(recency)
		ranked[i].CompletenessScore = round4(completeness)
		ranked[i].TotalScore = round4(total)
	}

	sort.SliceStable(ranked, func(i, j int) bool {
		if ranked[i].TotalScore != ranked[j].TotalScore {
			return ranked[i].TotalScore > ranked[j].TotalScore
		}
		// Tie breaker: recency then ID
		if ranked[i].PublicationYear != ranked[j].PublicationYear {
			return ranked[i].PublicationYear > ranked[j].PublicationYear
		}
		return ranked[i].ID < ranked[j].ID
	})

	if len(ranked) > limit {
		return ranked[:limit]
	}
	return ranked
}

func calculateRelevance(query string, tokens []string, a *models.Article) float64 {
	if len(tokens) == 0 {
		return 0.5
	}

	titleLower := strings.ToLower(a.Title)
	abstractLower := strings.ToLower(a.Abstract)
	queryLower := strings.TrimSpace(strings.ToLower(query))

	score := 0.0

	// Exact phrase bonus
	if strings.Contains(titleLower, queryLower) {
		score += 0.4
	} else if strings.Contains(abstractLower, queryLower) {
		score += 0.2
	}

	// Token overlap
	matchedTokens := 0
	for _, t := range tokens {
		if strings.Contains(titleLower, t) {
			score += 0.3 / float64(len(tokens))
			matchedTokens++
		} else if strings.Contains(abstractLower, t) {
			score += 0.15 / float64(len(tokens))
			matchedTokens++
		}
	}

	// Proportion bonus
	tokenRatio := float64(matchedTokens) / float64(len(tokens))
	score += 0.3 * tokenRatio

	if score > 1.0 {
		return 1.0
	}
	if score < 0.1 && len(a.Title) > 0 {
		return 0.1
	}
	return score
}

func calculateRecency(pubYear int, currentYear int) float64 {
	if pubYear <= 0 {
		return 0.3
	}
	diff := currentYear - pubYear
	switch {
	case diff <= 0:
		return 1.0 // Current or future release year
	case diff == 1:
		return 0.95
	case diff == 2:
		return 0.88
	case diff == 3:
		return 0.80
	case diff == 4:
		return 0.70
	case diff <= 6:
		return 0.55
	case diff <= 10:
		return 0.40
	default:
		return 0.20
	}
}

func calculateCompleteness(a *models.Article) float64 {
	score := 0.0

	if a.DOI != "" {
		score += 0.30
	}
	if len(a.Abstract) >= 120 {
		score += 0.30
	} else if len(a.Abstract) > 0 {
		score += 0.15
	}
	if len(a.Authors) > 0 {
		score += 0.20
	}
	if a.Journal != "" {
		score += 0.10
	}
	if a.IsOpenAccess || a.PDFURL != "" {
		score += 0.10
	}

	if score > 1.0 {
		return 1.0
	}
	return score
}

func tokenize(s string) []string {
	words := strings.Fields(strings.ToLower(s))
	var tokens []string
	for _, w := range words {
		w = strings.Trim(w, `.,!?;:"'()[]{}`)
		if len(w) > 2 && !stopWords[w] {
			tokens = append(tokens, w)
		}
	}
	return tokens
}

func round4(val float64) float64 {
	return float64(int(val*10000+0.5)) / 10000
}
