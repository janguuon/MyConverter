package hwpx

import "encoding/xml"

// SettingsXML represents the settings.xml structure
type SettingsXML struct {
	XMLName      xml.Name     `xml:"hc:settings"`
	Config       ConfigType   `xml:"hc:config"`
	CaretState   CaretState   `xml:"hc:caret-state"`
	ViewSettings ViewSettings `xml:"hc:view-settings"`
}

// ConfigType represents document configuration
type ConfigType struct {
	ZoomPercent    int    `xml:"hc:zoom-percent,attr"`
	ShowGutter     bool   `xml:"hc:show-gutter,attr"`
	ShowParaMarks  bool   `xml:"hc:show-para-marks,attr"`
	ViewMode       string `xml:"hc:view-mode,attr"`
}

// CaretState represents cursor position
type CaretState struct {
	List []CaretPosition `xml:"hc:list"`
}

// CaretPosition represents a cursor position
type CaretPosition struct {
	PageNum    int `xml:"hc:page-num,attr"`
	LineNum    int `xml:"hc:line-num,attr"`
	ColumnNum  int `xml:"hc:column-num,attr"`
}

// ViewSettings represents view settings
type ViewSettings struct {
	ShowLineNumber bool `xml:"hc:show-line-number,attr"`
	ShowRuler      bool `xml:"hc:show-ruler,attr"`
}
