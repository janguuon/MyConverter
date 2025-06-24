package readers

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// HWPXContent represents the extracted text content from HWPX
type HWPXContent struct {
	Text []string
}

// ExtractHWPXContent extracts text content from section0.xml in HWPX file
func ExtractHWPXContent(zipReader *zip.ReadCloser) (*HWPXContent, error) {
	var content HWPXContent

	// Find and read section0.xml
	for _, file := range zipReader.File {
		if strings.HasSuffix(file.Name, "section0.xml") {
			rc, err := file.Open()
			if err != nil {
				return nil, fmt.Errorf("failed to open section0.xml: %v", err)
			}
			defer rc.Close()

			// Read the XML content
			xmlContent, err := io.ReadAll(rc)
			if err != nil {
				return nil, fmt.Errorf("failed to read XML content: %v", err)
			}

			// Parse XML to extract text from hp:t elements
			decoder := xml.NewDecoder(strings.NewReader(string(xmlContent)))
			var texts []string
			var isHPT bool

			for {
				token, err := decoder.Token()
				if err == io.EOF {
					break
				}
				if err != nil {
					return nil, fmt.Errorf("failed to parse XML: %v", err)
				}

				switch t := token.(type) {
				case xml.StartElement:
					if t.Name.Local == "t" {
						isHPT = true
					}
				case xml.CharData:
					if isHPT {
						text := strings.TrimSpace(string(t))
						if text != "" {
							texts = append(texts, text)
						}
					}
				case xml.EndElement:
					if t.Name.Local == "t" {
						isHPT = false
					}
				}
			}

			content.Text = texts
			break
		}
	}

	return &content, nil
}
