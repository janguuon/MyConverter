package writers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/signintech/gopdf"
)

// PDFWriter handles converting document content to PDF files
type PDFWriter struct {
	// Output directory for PDF files
	OutputDir string
	// Page settings
	PageSize *gopdf.Rect // e.g., A4 size
	FontSize int        // in points
}

// NewPDFWriter creates a new PDFWriter with default settings
func NewPDFWriter(outputDir string) *PDFWriter {
	return &PDFWriter{
		OutputDir: outputDir,
		PageSize:  gopdf.PageSizeA4,
		FontSize: 48,
	}
}

// WriteTexts renders text in the PDF
func (w *PDFWriter) WriteTexts(outputPath string, text string) error {
	// Split text into lines
	lines := strings.Split(text, "\n")

	// Create new PDF
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *w.PageSize})
	
	// Add a new page
	pdf.AddPage()

	// Add Korean font (using Malgun Gothic)
	err := pdf.AddTTFFont("malgun", "C:\\Windows\\Fonts\\malgun.ttf")
	if err != nil {
		return fmt.Errorf("failed to load font: %v", err)
	}

	// Set font
	err = pdf.SetFont("malgun", "", 10) // 10pt font size
	if err != nil {
		return fmt.Errorf("failed to set font: %v", err)
	}

	// Start from top with margin
	margin := 50.0
	lineHeight := 14.0 // 1.4x line spacing for 10pt font
	y := margin

	// Draw each line of text
	for _, line := range lines {
		pdf.SetX(margin)
		pdf.SetY(y)
		pdf.Cell(nil, line)
		y += lineHeight

		// Add new page if we're near the bottom
		if y > w.PageSize.H-margin {
			pdf.AddPage()
			y = margin
		}
	}

	// Ensure output directory exists
	if err := os.MkdirAll(w.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Save the PDF
	outputPath = filepath.Join(w.OutputDir, filepath.Base(outputPath))
	err = pdf.WritePdf(outputPath)
	if err != nil {
		return fmt.Errorf("failed to save PDF: %v", err)
	}

	return nil
}

// Write creates a PDF with sample text and saves it
func (w *PDFWriter) Write(outputPath string) error {
	return w.WriteTexts(outputPath, "Sample Text")
}

// UpdateMetadata updates the metadata of an existing PDF file
func (w *PDFWriter) UpdateMetadata(outputPath string, metadata map[string]string) error {
	// Create new PDF
	pdf := gopdf.GoPdf{}

	// Set metadata
	if title, ok := metadata["Title"]; ok {
		pdf.SetInfo(gopdf.PdfInfo{
			Title:    title,
			Author:   metadata["Author"],
			Creator:  metadata["Creator"],
			Subject:  metadata["Subject"],
			Producer: metadata["Producer"],
		})
	}

	return nil
}
