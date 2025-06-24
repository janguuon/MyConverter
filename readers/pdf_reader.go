package readers

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
	"myconverter/interfaces"
)

type PDFReader struct {
	filePath string
	content  *interfaces.PDFContent
}

func NewPDFReader(filePath string) *PDFReader {
	return &PDFReader{
		filePath: filePath,
		content:  &interfaces.PDFContent{},
	}
}

func (r *PDFReader) Read(filePath string) error {
	r.filePath = filePath
	pdfContent, err := r.ReadPDF()
	if err != nil {
		return err
	}

	// 메타데이터와 텍스트 내용을 포함한 결과 생성
	var result strings.Builder

	// 메타데이터 정보 추가
	result.WriteString("[PDF Metadata]\n")
	for key, value := range pdfContent.Metadata {
		result.WriteString(fmt.Sprintf("%s: %s\n", key, value))
	}
	result.WriteString("\n")

	// 페이지별 텍스트 내용 추가
	result.WriteString("[PDF Content]\n")
	for _, page := range pdfContent.Pages {
		result.WriteString(fmt.Sprintf("=== Page %d ===\n", page.Number))
		result.WriteString(page.Text)
		result.WriteString("\n\n")
	}

	// 결과 저장
	r.content.Text = result.String()
	return nil
}

func (r *PDFReader) ReadPDF() (*interfaces.PDFContent, error) {
	// PDF 파일 열기
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open PDF file: %v", err)
	}
	defer file.Close()

	// PDF 문서 생성
	pdfReader, err := model.NewPdfReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to create PDF reader: %v", err)
	}

	// 메타데이터 추출
	r.content.Metadata = make(map[string]string)
	if info, err := pdfReader.GetPdfInfo(); err == nil {
		if info.Title != nil {
			r.content.Metadata["Title"] = info.Title.String()
		}
		if info.Author != nil {
			r.content.Metadata["Author"] = info.Author.String()
		}
		if info.Subject != nil {
			r.content.Metadata["Subject"] = info.Subject.String()
		}
		if info.Creator != nil {
			r.content.Metadata["Creator"] = info.Creator.String()
		}
		if info.Producer != nil {
			r.content.Metadata["Producer"] = info.Producer.String()
		}
	} else {
		// 기본 메타데이터 설정
		r.content.Metadata["Title"] = "Unknown"
		r.content.Metadata["Author"] = "Unknown"
		r.content.Metadata["Subject"] = "Unknown"
		r.content.Metadata["Creator"] = "MyConverter"
		r.content.Metadata["Producer"] = "MyConverter"
	}

	// 페이지 수 확인
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return nil, fmt.Errorf("failed to get page count: %v", err)
	}

	// 페이지별 텍스트 추출
	r.content.Pages = make([]interfaces.PDFPage, numPages)
	var textBuilder strings.Builder

	for i := 0; i < numPages; i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			return nil, fmt.Errorf("failed to get page %d: %v", i+1, err)
		}

		// 페이지에서 텍스트 추출
		extractor, err := extractor.New(page)
		if err != nil {
			return nil, fmt.Errorf("failed to create text extractor for page %d: %v", i+1, err)
		}

		text, err := extractor.ExtractText()
		if err != nil {
			return nil, fmt.Errorf("failed to extract text from page %d: %v", i+1, err)
		}

		textBuilder.WriteString(text)
		textBuilder.WriteString("\n")

		r.content.Pages[i] = interfaces.PDFPage{
			Number: i + 1,
			Text:   text,
			Images: nil,
		}
	}

	r.content.Text = textBuilder.String()
	return r.content, nil
}

func (r *PDFReader) ReadStream() (io.ReadCloser, error) {
	return os.Open(r.filePath)
}

func (r *PDFReader) GetContent() *interfaces.PDFContent {
	return r.content
}
