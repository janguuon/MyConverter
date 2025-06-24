package hwpx

import "encoding/xml"

// ManifestXML represents the META-INF/manifest.xml structure
type ManifestXML struct {
	XMLName   xml.Name        `xml:"manifest"`
	FileEntry []FileEntryType `xml:"file-entry"`
}

// FileEntryType represents a file entry in manifest.xml
type FileEntryType struct {
	FullPath      string `xml:"full-path,attr"`
	MediaType     string `xml:"media-type,attr"`
	Size          int64  `xml:"size,attr,omitempty"`
	LastModified  string `xml:"last-modified,attr,omitempty"`
}
