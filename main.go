package main

import (
	"archive/zip"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"myconverter/interfaces"
	"myconverter/readers"
	"myconverter/writers"
)

// GetFileReader returns appropriate reader based on file extension
func GetFileReader(filePath string) interfaces.FileReader {
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".hwp":
		return &readers.CFBReader{}
	case ".hwpx", ".zip":
		return &readers.ZipReader{}
	case ".pdf":
		return readers.NewPDFReader(filePath)
	default:
		return nil
	}
}

// processReader handles the reading and display of file contents
func processReader(reader interfaces.FileReader, filePath string) error {
	err := reader.Read(filePath)
	if err != nil {
		return err
	}

	// Type assert to access specific reader methods
	switch r := reader.(type) {
	case *readers.CFBReader:
		// Display CFB entries
		for _, entry := range r.GetEntries() {
			fmt.Printf("Entry: %s, Size: %d bytes\n", entry.Name, entry.Size)
		}

	case *readers.ZipReader:
		// Display ZIP entries
		if strings.HasSuffix(strings.ToLower(filePath), ".hwpx") {
			// For HWPX files, show only XML contents
			for _, file := range r.GetXMLFiles() {
				fmt.Printf("\nXML File: %s, Size: %d bytes\n", file.Name, file.Size)
				if len(file.Content) > 500 {
					fmt.Printf("Content (first 500 bytes):\n%s...\n", file.Content[:500])
				} else {
					fmt.Printf("Content:\n%s\n", file.Content)
				}
			}
		} else {
			// For regular ZIP files, show all contents
			for _, file := range r.GetFiles() {
				fmt.Printf("File: %s, Size: %d bytes\n", file.Name, file.Size)
			}
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  Read:    myconverter read <filepath>")
		fmt.Println("  Write:   myconverter write <output.zip> <file1> [file2] [dir1] ...")
		fmt.Println("  Convert: myconverter convert <input_file> <output.(png|pdf)>")
		return
	}

	command := strings.ToLower(os.Args[1])

	switch command {
	case "read":
		filePath := os.Args[2]
		reader := GetFileReader(filePath)
		if reader == nil {
			fmt.Println("Unsupported file format")
			return
		}

		err := processReader(reader, filePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

	case "write":
		if len(os.Args) < 4 {
			fmt.Println("Error: Please provide output zip file and at least one source file/directory")
			return
		}

		writer := &writers.ZipWriter{
			TargetPath: os.Args[2],
			Sources:    os.Args[3:],
		}

		err := writer.Write(writer.TargetPath)
		if err != nil {
			fmt.Printf("Error creating zip file: %v\n", err)
			return
		}
		fmt.Printf("Successfully created zip file: %s\n", writer.TargetPath)

	case "convert":
		if len(os.Args) < 4 {
			fmt.Println("Error: Please provide input file and output path")
			return
		}

		inputFile := os.Args[2]
		outputFile := os.Args[3]
		outputDir := filepath.Dir(outputFile)
		ext := strings.ToLower(filepath.Ext(outputFile))

		// Open HWPX file
		zipReader, err := zip.OpenReader(inputFile)
		if err != nil {
			fmt.Printf("Error opening HWPX file: %v\n", err)
			return
		}
		defer zipReader.Close()

		// Get appropriate reader based on input file extension
		reader := GetFileReader(inputFile)
		if reader == nil {
			fmt.Println("Unsupported input file format")
			return
		}

		// Read content based on file type
		var content *interfaces.PDFContent

		switch r := reader.(type) {
		case *readers.PDFReader:
			// Read PDF content
			content, err = r.ReadPDF()
			if err != nil {
				fmt.Printf("Error reading PDF content: %v\n", err)
				return
			}

		case *readers.ZipReader:
			// For HWPX files
			if strings.HasSuffix(strings.ToLower(inputFile), ".hwpx") {
				// Extract text content from HWPX
				hwpxContent, err := readers.ExtractHWPXContent(zipReader)
				if err != nil {
					fmt.Printf("Error extracting HWPX content: %v\n", err)
					return
				}
				// Convert HWPX content to PDFContent format
				content = &interfaces.PDFContent{
					Text: strings.Join(hwpxContent.Text, "\n"),
					Pages: []interfaces.PDFPage{{
						Number: 1,
						Text:   strings.Join(hwpxContent.Text, "\n"),
					}},
				}
			}

		default:
			fmt.Printf("Unsupported input file type\n")
			return
		}

		switch ext {
		case ".png":
			// Create image writer
			imageWriter, err := writers.NewImageWriter(outputDir)
			if err != nil {
				fmt.Printf("Error creating image writer: %v\n", err)
				return
			}

			// Create image with text
			err = imageWriter.WriteTexts(outputFile, content.Text)
			if err != nil {
				fmt.Printf("Error creating image: %v\n", err)
				return
			}

			fmt.Printf("Successfully created image file: %s\n", outputFile)

		case ".pdf":
			// Create PDF writer
			pdfWriter := writers.NewPDFWriter(outputDir)

			// Create PDF with text
			err := pdfWriter.WriteTexts(outputFile, content.Text)
			if err != nil {
				fmt.Printf("Error creating PDF: %v\n", err)
				return
			}

			// If input was PDF and output is PDF, copy metadata
			if inputPDF, ok := reader.(*readers.PDFReader); ok {
				if pdfContent := inputPDF.GetContent(); pdfContent != nil && pdfContent.Metadata != nil {
					err = pdfWriter.UpdateMetadata(outputFile, pdfContent.Metadata)
					if err != nil {
						fmt.Printf("Warning: Failed to copy PDF metadata: %v\n", err)
					}
				}
			}
			if err != nil {
				fmt.Printf("Error creating PDF: %v\n", err)
				return
			}

			fmt.Printf("Successfully created PDF file: %s\n", outputFile)

		case ".txt":
			// Create text writer
			textWriter := writers.NewTextWriter(outputDir)
			err = textWriter.WriteTexts(outputFile, content.Text)
			if err != nil {
				fmt.Printf("Error creating text file: %v\n", err)
				return
			}

			fmt.Printf("Successfully created text file: %s\n", outputFile)

		default:
			fmt.Printf("Unsupported output format: %s\n", ext)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}
