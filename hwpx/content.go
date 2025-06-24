package hwpx

import "encoding/xml"

// ContentXML represents the content/content.xml structure
type ContentXML struct {
	XMLName     xml.Name    `xml:"hml"`
	Head        HeadType    `xml:"head"`
	Body        BodyType    `xml:"body"`
	Namespace   string      `xml:"xmlns,attr"`
	Version     string      `xml:"version,attr"`
	SecElemInfo SecElemInfo `xml:"sec-elem-info,omitempty"`
}

// HeadType represents document head information
type HeadType struct {
	MetaData   MetaDataType   `xml:"metadata"`
	Properties PropertiesType `xml:"properties"`
	Style      StylesType     `xml:"style"`
	References ReferencesType `xml:"references"`
}

// MetaDataType represents document metadata
type MetaDataType struct {
	Title       string `xml:"title"`
	Subject     string `xml:"subject"`
	Author      string `xml:"author"`
	Date        string `xml:"date"`
	Keywords    string `xml:"keywords"`
	Description string `xml:"description"`
	Generator   string `xml:"generator"`
}

// PropertiesType represents document properties
type PropertiesType struct {
	PageDef      PageDefType      `xml:"page-def"`
	FootnotesDef FootnotesDefType `xml:"footnotes-def"`
	EndnotesDef  EndnotesDefType  `xml:"endnotes-def"`
	AutoNumDef   AutoNumDefType   `xml:"auto-num-def"`
}

// PageDefType represents page definition
type PageDefType struct {
	Width     int     `xml:"width,attr"`
	Height    int     `xml:"height,attr"`
	Landscape bool    `xml:"landscape,attr"`
	Margins   Margins `xml:"margins"`
}

// Margins represents page margins
type Margins struct {
	Left   int `xml:"left,attr"`
	Right  int `xml:"right,attr"`
	Top    int `xml:"top,attr"`
	Bottom int `xml:"bottom,attr"`
}

// StylesType represents document styles
type StylesType struct {
	ParaStyles  []ParaStyle  `xml:"para-style"`
	CharStyles  []CharStyle  `xml:"char-style"`
	TableStyles []TableStyle `xml:"table-style"`
}

// ParaStyle represents paragraph style
type ParaStyle struct {
	ID          string  `xml:"id,attr"`
	Name        string  `xml:"name,attr"`
	Align       string  `xml:"align,attr,omitempty"`
	LineHeight  float64 `xml:"line-height,attr,omitempty"`
	MarginLeft  int     `xml:"margin-left,attr,omitempty"`
	MarginRight int     `xml:"margin-right,attr,omitempty"`
}

// CharStyle represents character style
type CharStyle struct {
	ID       string `xml:"id,attr"`
	Name     string `xml:"name,attr"`
	FontFace string `xml:"font-face,attr,omitempty"`
	FontSize int    `xml:"font-size,attr,omitempty"`
	Bold     bool   `xml:"bold,attr,omitempty"`
	Italic   bool   `xml:"italic,attr,omitempty"`
}

// TableStyle represents table style
type TableStyle struct {
	ID          string `xml:"id,attr"`
	Name        string `xml:"name,attr"`
	BorderType  string `xml:"border-type,attr,omitempty"`
	CellSpacing int    `xml:"cell-spacing,attr,omitempty"`
}

// ReferencesType represents document references
type ReferencesType struct {
	Fonts       []FontType      `xml:"font"`
	BorderTypes []BorderType    `xml:"border-type"`
	CharShapes  []CharShapeType `xml:"char-shape"`
	TabDefs     []TabDefType    `xml:"tab-def"`
}

// FontType represents font information
type FontType struct {
	ID       string `xml:"id,attr"`
	Name     string `xml:"name,attr"`
	Type     string `xml:"type,attr"`
	IsEmbedd bool   `xml:"is-embedd,attr,omitempty"`
}

// BorderType represents border information
type BorderType struct {
	ID    string `xml:"id,attr"`
	Type  string `xml:"type,attr"`
	Width int    `xml:"width,attr"`
	Color string `xml:"color,attr"`
	Style string `xml:"style,attr"`
}

// BodyType represents document body
type BodyType struct {
	SectionDef SectionDefType `xml:"section-def"`
	Sections   []SectionType  `xml:"section"`
}

// SectionDefType represents section definition
type SectionDefType struct {
	ID            string `xml:"id,attr"`
	PageWidth     int    `xml:"page-width,attr"`
	PageHeight    int    `xml:"page-height,attr"`
	PageMarginTop int    `xml:"page-margin-top,attr"`
}

// SectionType represents a section
type SectionType struct {
	ID         string          `xml:"id,attr"`
	Paragraphs []ParagraphType `xml:"p"`
	Tables     []TableType     `xml:"table"`
}

// ParagraphType represents a paragraph
type ParagraphType struct {
	ID      string    `xml:"id,attr"`
	StyleID string    `xml:"style-id,attr,omitempty"`
	Runs    []RunType `xml:"run"`
}

// RunType represents a text run
type RunType struct {
	StyleID string `xml:"style-id,attr,omitempty"`
	Text    string `xml:"text"`
}

// TableType represents a table
type TableType struct {
	ID       string    `xml:"id,attr"`
	RowsType []RowType `xml:"row"`
	Cols     int       `xml:"cols,attr"`
	Rows     int       `xml:"rows,attr"`
}

// RowType represents a table row
type RowType struct {
	Cells []CellType `xml:"cell"`
}

// CellType represents a table cell
type CellType struct {
	ColSpan int             `xml:"col-span,attr,omitempty"`
	RowSpan int             `xml:"row-span,attr,omitempty"`
	Content []ParagraphType `xml:"p"`
}
