package writers

import (
	"fmt"
	"os"
	"path/filepath"
)

// TextWriter handles writing plain text files.
type TextWriter struct {
	// Output directory for text files
	OutputDir string
}

// NewTextWriter creates a new TextWriter.
func NewTextWriter(outputDir string) *TextWriter {
	return &TextWriter{OutputDir: outputDir}
}

// WriteTexts saves the given text to a file.
func (w *TextWriter) WriteTexts(outputPath string, text string) error {
	if err := os.MkdirAll(w.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	outputPath = filepath.Join(w.OutputDir, filepath.Base(outputPath))
	return os.WriteFile(outputPath, []byte(text), 0644)
}

// Write writes sample text to the specified file.
func (w *TextWriter) Write(outputPath string) error {
	return w.WriteTexts(outputPath, "Sample Text")
}
