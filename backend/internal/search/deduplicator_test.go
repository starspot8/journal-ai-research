package search

import (
	"testing"

	"journal-ai-research/backend/internal/models"
)

func TestDeduplicateArticles(t *testing.T) {
	articles := []models.Article{
		{
			ID:           "pmid-12345",
			Title:        "Microbiome Diversity and Human Health",
			Abstract:     "Short abstract.",
			DOI:          "10.1038/s41586-020-0001",
			PMID:         "12345",
			Source:       "pubmed",
			IsOpenAccess: false,
		},
		{
			ID:           "epmc-12345",
			Title:        "Microbiome Diversity and Human Health.",
			Abstract:     "A much more detailed and longer abstract describing the gut microbiome and host interactions.",
			DOI:          "10.1038/s41586-020-0001", // Identical DOI
			PMCID:        "PMC99999",
			Source:       "europepmc",
			IsOpenAccess: true,
			PDFURL:       "https://europepmc.org/pdf/PMC99999.pdf",
		},
		{
			ID:      "crossref-abc",
			Title:   "Unrelated Study on Deep Learning in Robotics",
			DOI:     "10.1109/robot.2023.01",
			Source:  "crossref",
			Journal: "IEEE Robotics",
		},
	}

	deduped := DeduplicateArticles(articles)

	if len(deduped) != 2 {
		t.Fatalf("expected 2 unique articles after deduplication, got %d", len(deduped))
	}

	merged := deduped[0]
	if merged.PMID != "12345" || merged.PMCID != "PMC99999" {
		t.Errorf("expected merged identifiers (PMID: %s, PMCID: %s)", merged.PMID, merged.PMCID)
	}
	if !merged.IsOpenAccess {
		t.Errorf("expected merged article to inherit Open Access status")
	}
	if len(merged.Abstract) <= len("Short abstract.") {
		t.Errorf("expected merged article to keep the more comprehensive abstract")
	}
}
