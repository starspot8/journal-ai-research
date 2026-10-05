package pdf

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
	"journal-ai-research/backend/internal/models"
)

// GenerateResearchReport creates a PDF document summarizing research query and selected articles
func GenerateResearchReport(query string, articles []models.Article) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(15, 15, 15)
	pdf.AddPage()

	// Header
	pdf.SetFont("Helvetica", "B", 18)
	pdf.SetTextColor(30, 64, 175) // Primary Blue (#1E40AF)
	pdf.Cell(0, 10, "Journal AI Research")
	pdf.Ln(7)

	pdf.SetFont("Helvetica", "", 11)
	pdf.SetTextColor(107, 114, 128) // Text Secondary (#6B7280)
	pdf.Cell(0, 7, "Scientific Literature Research Report")
	pdf.Ln(10)

	// Divider line
	pdf.SetDrawColor(229, 231, 235) // Border (#E5E7EB)
	pdf.SetLineWidth(0.5)
	pdf.Line(15, pdf.GetY(), 195, pdf.GetY())
	pdf.Ln(5)

	// Query Summary Box
	pdf.SetFillColor(249, 250, 251) // Surface Card (#F9FAFB)
	pdf.Rect(15, pdf.GetY(), 180, 22, "F")

	pdf.SetFont("Helvetica", "B", 10)
	pdf.SetTextColor(17, 24, 39) // Text Primary (#111827)
	pdf.Text(18, pdf.GetY()+6, "Query:")

	pdf.SetFont("Helvetica", "", 10)
	pdf.SetTextColor(30, 64, 175)
	pdf.Text(35, pdf.GetY()+6, query)

	pdf.SetFont("Helvetica", "B", 10)
	pdf.SetTextColor(17, 24, 39)
	pdf.Text(18, pdf.GetY()+13, "Generated At:")

	pdf.SetFont("Helvetica", "", 10)
	pdf.SetTextColor(107, 114, 128)
	pdf.Text(45, pdf.GetY()+13, time.Now().UTC().Format("02 January 2006, 15:04 UTC"))

	pdf.SetFont("Helvetica", "B", 10)
	pdf.SetTextColor(17, 24, 39)
	pdf.Text(18, pdf.GetY()+19, "Total Articles:")

	pdf.SetFont("Helvetica", "", 10)
	pdf.SetTextColor(107, 114, 128)
	pdf.Text(45, pdf.GetY()+19, fmt.Sprintf("%d articles", len(articles)))

	pdf.Ln(28)

	// Articles Listing
	for i, a := range articles {
		// Check for page break if close to bottom
		if pdf.GetY() > 240 {
			pdf.AddPage()
		}

		// Article Title
		pdf.SetFont("Helvetica", "B", 12)
		pdf.SetTextColor(17, 24, 39)
		titleText := fmt.Sprintf("%d. %s", i+1, a.Title)
		pdf.MultiCell(180, 6, titleText, "", "L", false)
		pdf.Ln(1)

		// Metadata (Authors, Journal, Year)
		var authorNames []string
		for _, auth := range a.Authors {
			authorNames = append(authorNames, auth.Name)
		}
		authorStr := strings.Join(authorNames, ", ")
		if len(authorStr) > 100 {
			authorStr = authorStr[:97] + "..."
		}
		if authorStr == "" {
			authorStr = "Authors not indexed"
		}

		pdf.SetFont("Helvetica", "I", 9)
		pdf.SetTextColor(107, 114, 128)
		metaLine := fmt.Sprintf("%s | %s (%d)", authorStr, a.Journal, a.PublicationYear)
		pdf.MultiCell(180, 5, metaLine, "", "L", false)

		// Identifiers
		var idParts []string
		if a.DOI != "" {
			idParts = append(idParts, "DOI: "+a.DOI)
		}
		if a.PMID != "" {
			idParts = append(idParts, "PMID: "+a.PMID)
		}
		if a.PMCID != "" {
			idParts = append(idParts, "PMCID: "+a.PMCID)
		}
		if a.IsOpenAccess {
			idParts = append(idParts, "[Open Access]")
		}
		if len(idParts) > 0 {
			pdf.SetFont("Helvetica", "", 8)
			pdf.SetTextColor(14, 165, 233) // Secondary blue (#0EA5E9)
			pdf.Cell(0, 5, strings.Join(idParts, " | "))
			pdf.Ln(5)
		}

		// Abstract
		if a.Abstract != "" {
			abstractText := a.Abstract
			if len(abstractText) > 400 {
				abstractText = abstractText[:397] + "..."
			}
			pdf.SetFont("Helvetica", "", 9)
			pdf.SetTextColor(55, 65, 81)
			pdf.MultiCell(180, 4.5, abstractText, "", "J", false)
			pdf.Ln(2)
		}

		// URL Link
		if a.URL != "" {
			pdf.SetFont("Helvetica", "U", 8)
			pdf.SetTextColor(30, 64, 175)
			pdf.Cell(0, 4, "Original Article: "+a.URL)
			pdf.Ln(5)
		}

		// Separator
		pdf.SetDrawColor(243, 244, 246)
		pdf.Line(15, pdf.GetY(), 195, pdf.GetY())
		pdf.Ln(4)
	}

	// Legal Disclaimer Footer
	if pdf.GetY() > 255 {
		pdf.AddPage()
	}
	pdf.Ln(4)
	pdf.SetFont("Helvetica", "I", 8)
	pdf.SetTextColor(156, 163, 175)
	disclaimer := "Disclaimer: Journal AI Research is a scientific literature search and research assistance platform. " +
		"Metadata and abstracts are retrieved from PubMed, Europe PMC, and Crossref. " +
		"This platform is not intended to provide medical diagnosis, treatment recommendations, or individualized medical advice."
	pdf.MultiCell(180, 3.5, disclaimer, "", "C", false)

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, fmt.Errorf("failed to generate PDF output: %w", err)
	}

	return buf.Bytes(), nil
}
