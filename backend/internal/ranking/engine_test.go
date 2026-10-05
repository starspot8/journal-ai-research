package ranking

import (
	"testing"
	"time"

	"journal-ai-research/backend/internal/models"
)

func TestRankArticles(t *testing.T) {
	currentYear := time.Now().Year()

	articles := []models.Article{
		{
			ID:              "art-1",
			Title:           "Artificial Intelligence in Glaucoma Detection",
			Abstract:        "Comprehensive study on deep learning applications for glaucoma screening with high accuracy.",
			Journal:         "Ophthalmology",
			PublicationYear: currentYear,
			DOI:             "10.1016/j.ophtha.2024.01.001",
			Authors:         []models.Author{{Name: "Dr. Alice Smith"}},
			IsOpenAccess:    true,
		},
		{
			ID:              "art-2",
			Title:           "Cataract Surgery Historical Review",
			Abstract:        "Historical aspects of cataract extraction in the 1990s.",
			Journal:         "Eye History",
			PublicationYear: currentYear - 15,
			DOI:             "10.1000/182",
		},
		{
			ID:              "art-3",
			Title:           "AI Applications in Eye Disease Detection",
			Abstract:        "Recent advances in machine learning for retinal disease and glaucoma.",
			Journal:         "Nature Medicine",
			PublicationYear: currentYear - 1,
			DOI:             "10.1038/s41591-023-0001",
			Authors:         []models.Author{{Name: "Dr. Bob Lee"}},
			IsOpenAccess:    true,
		},
	}

	ranked := RankArticles("artificial intelligence glaucoma detection", articles, 10)

	if len(ranked) != 3 {
		t.Fatalf("expected 3 ranked articles, got %d", len(ranked))
	}

	// First article should be art-1 because of exact title matches, recency, and metadata completeness
	if ranked[0].ID != "art-1" {
		t.Errorf("expected top article to be art-1, got %s (score: %f)", ranked[0].ID, ranked[0].TotalScore)
	}

	for i := range ranked {
		if ranked[i].TotalScore <= 0 {
			t.Errorf("article %s has non-positive score %f", ranked[i].ID, ranked[i].TotalScore)
		}
	}
}
