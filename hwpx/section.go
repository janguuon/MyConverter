package hwpx

import "encoding/xml"

// SectionXML represents the content/section0.xml structure
type SectionXML struct {
	XMLName   xml.Name      `xml:"section"`
	ID        string        `xml:"id,attr"`
	Type      string        `xml:"type,attr"`
	Paragraphs []ParagraphType `xml:"p"`
	Tables    []TableType     `xml:"table"`
	Headers   []HeaderType    `xml:"header"`
	Footers   []FooterType    `xml:"footer"`
}

// HeaderType represents a header
type HeaderType struct {
	ID         string        `xml:"id,attr"`
	Type       string        `xml:"type,attr"`
	Paragraphs []ParagraphType `xml:"p"`
}

// FooterType represents a footer
type FooterType struct {
	ID         string        `xml:"id,attr"`
	Type       string        `xml:"type,attr"`
	Paragraphs []ParagraphType `xml:"p"`
}

// FootnotesDefType represents footnotes definition
type FootnotesDefType struct {
	AutoNumFormat   string `xml:"auto-num-format,attr"`
	AutoNumStart    int    `xml:"auto-num-start,attr"`
	PlaceEndOfPage  bool   `xml:"place-end-of-page,attr"`
}

// EndnotesDefType represents endnotes definition
type EndnotesDefType struct {
	AutoNumFormat  string `xml:"auto-num-format,attr"`
	AutoNumStart   int    `xml:"auto-num-start,attr"`
}

// AutoNumDefType represents auto numbering definition
type AutoNumDefType struct {
	Type     string `xml:"type,attr"`
	Format   string `xml:"format,attr"`
	Start    int    `xml:"start,attr"`
}

// CharShapeType represents character shape information
type CharShapeType struct {
	ID       string `xml:"id,attr"`
	FontID   string `xml:"font-id,attr"`
	FontSize int    `xml:"font-size,attr"`
	Bold     bool   `xml:"bold,attr"`
	Italic   bool   `xml:"italic,attr"`
}

// TabDefType represents tab definition
type TabDefType struct {
	ID      string `xml:"id,attr"`
	Type    string `xml:"type,attr"`
	Leader  string `xml:"leader,attr"`
	Position int   `xml:"position,attr"`
}

// SecElemInfo represents section element information
type SecElemInfo struct {
	SecPr []SecPrType `xml:"sec-pr"`
}

// SecPrType represents section properties
type SecPrType struct {
	ID            string `xml:"id,attr"`
	PageWidth     int    `xml:"page-width,attr"`
	PageHeight    int    `xml:"page-height,attr"`
	PageMarginTop int    `xml:"page-margin-top,attr"`
}
