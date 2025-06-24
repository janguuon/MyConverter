package readers

import (
	"archive/zip"
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

// ZipReader implements FileReader for ZIP format
type ZipReader struct {
	FilePath string
	Files    []ZipEntry
}

// ZipEntry represents a file in the ZIP archive
type ZipEntry struct {
	Name     string
	Size     uint64
	Content  []byte
	IsXML    bool
}

// Read implements ZIP file reading
func (r *ZipReader) Read(filePath string) error {
	r.FilePath = filePath
	reader, err := zip.OpenReader(filePath)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %v", err)
	}
	defer reader.Close()

	ext := strings.ToLower(filepath.Ext(filePath))
	isHwpx := ext == ".hwpx"

	// Read all files in the archive
	for _, file := range reader.File {
		// For .hwpx files, only process XML files
		if isHwpx && !strings.HasSuffix(strings.ToLower(file.Name), ".xml") {
			continue
		}

		// Open the file inside the zip
		rc, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open file inside zip: %v", err)
		}

		// Read the content
		content, err := io.ReadAll(rc)
		if err != nil {
			rc.Close()
			return fmt.Errorf("failed to read file content: %v", err)
		}
		rc.Close()

		r.Files = append(r.Files, ZipEntry{
			Name:    file.Name,
			Size:    file.UncompressedSize64,
			Content: content,
			IsXML:   strings.HasSuffix(strings.ToLower(file.Name), ".xml"),
		})
	}

	return nil
}

// GetFiles returns all files in the ZIP archive
func (r *ZipReader) GetFiles() []ZipEntry {
	return r.Files
}

// GetXMLFiles returns only XML files from the ZIP archive
func (r *ZipReader) GetXMLFiles() []ZipEntry {
	var xmlFiles []ZipEntry
	for _, file := range r.Files {
		if file.IsXML {
			xmlFiles = append(xmlFiles, file)
		}
	}
	return xmlFiles
}
