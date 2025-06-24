package writers

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// ZipWriter handles ZIP file creation
type ZipWriter struct {
	TargetPath string
	Sources    []string
}

// Write implements FileWriter interface
func (w *ZipWriter) Write(filePath string) error {
	w.TargetPath = filePath
	return w.CreateZip()
}

// CreateZip creates a new ZIP file from the given sources
func (w *ZipWriter) CreateZip() error {
	// Create a new ZIP file
	zipFile, err := os.Create(w.TargetPath)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %v", err)
	}
	defer zipFile.Close()

	// Create a new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Process each source
	for _, source := range w.Sources {
		// Check if source exists
		sourceInfo, err := os.Stat(source)
		if err != nil {
			return fmt.Errorf("failed to stat %s: %v", source, err)
		}

		if sourceInfo.IsDir() {
			// Handle directory
			baseDir := filepath.Base(source)
			err = filepath.WalkDir(source, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				// Skip if it's the root directory
				if path == source {
					return nil
				}

				// Create zip entry path relative to the base directory
				relPath, err := filepath.Rel(source, path)
				if err != nil {
					return fmt.Errorf("failed to get relative path: %v", err)
				}
				zipPath := filepath.Join(baseDir, relPath)

				if d.IsDir() {
					// Create empty directory entry
					_, err = zipWriter.Create(zipPath + "/")
					return err
				}

				return w.addFileToZip(zipWriter, path, zipPath)
			})
		} else {
			// Handle single file
			err = w.addFileToZip(zipWriter, source, filepath.Base(source))
		}

		if err != nil {
			return fmt.Errorf("failed to add %s to zip: %v", source, err)
		}
	}

	return nil
}

// addFileToZip adds a single file to the zip archive
func (w *ZipWriter) addFileToZip(zipWriter *zip.Writer, sourcePath, zipPath string) error {
	// Open the source file
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Create the file in the zip archive
	writer, err := zipWriter.Create(zipPath)
	if err != nil {
		return err
	}

	// Copy the contents to the zip archive
	_, err = io.Copy(writer, sourceFile)
	return err
}
