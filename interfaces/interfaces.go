package interfaces

import "io"

// FileReader is the interface that wraps the basic Read method.
type FileReader interface {
	Read(filePath string) error
}

// PDFContent represents the content extracted from a PDF file
type PDFContent struct {
	Text     string
	Pages    []PDFPage
	Metadata map[string]string
}

// PDFPage represents a single page in a PDF document
type PDFPage struct {
	Number int
	Text   string
	Images [][]byte
}

// PDFReader is the interface that wraps the basic PDF reading methods
type PDFReader interface {
	ReadPDF() (*PDFContent, error)
	ReadStream() (io.ReadCloser, error)
}
