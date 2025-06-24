package writers

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

// ImageWriter handles converting document content to image files
type ImageWriter struct {
	// Width and Height of the output image in pixels
	Width  int
	Height int
	// Output directory for image files
	OutputDir string
	// Font settings
	font     *truetype.Font
	fontSize float64
}

// NewImageWriter creates a new ImageWriter with default settings
func NewImageWriter(outputDir string) (*ImageWriter, error) {
	// Load the font file (using NanumGothic as an example)
	fontBytes, err := ioutil.ReadFile("C:\\Windows\\Fonts\\malgun.ttf")
	if err != nil {
		return nil, fmt.Errorf("failed to load font file: %v", err)
	}

	// Parse the font
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font: %v", err)
	}

	return &ImageWriter{
		Width:     1920,    // Default width for A4 at 300 DPI
		Height:    2700,    // Default height for A4 at 300 DPI
		OutputDir: outputDir,
		font:     f,
		fontSize: 48,      // Default font size
	}, nil
}

// CreateImage creates a blank image with the specified dimensions
func (w *ImageWriter) CreateImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, w.Width, w.Height))

	// Fill with white background
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			img.Set(x, y, color.White)
		}
	}

	return img
}

// SaveImage saves the image to a PNG file
func (w *ImageWriter) SaveImage(img *image.RGBA, filename string) error {
	// Ensure output directory exists
	if err := os.MkdirAll(w.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Create output file
	outputPath := filepath.Join(w.OutputDir, filename)
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create image file: %v", err)
	}
	defer f.Close()

	// Encode and save as PNG
	if err := png.Encode(f, img); err != nil {
		return fmt.Errorf("failed to encode image: %v", err)
	}

	return nil
}

// WriteTexts renders text on the image and saves it
func (w *ImageWriter) WriteTexts(outputPath string, text string) error {
	// Split text into lines
	texts := strings.Split(text, "\n")
	// Create blank image
	img := w.CreateImage()

	// Create font context
	c := freetype.NewContext()
	c.SetDPI(300)
	c.SetFont(w.font)
	c.SetFontSize(10.0) // 10pt font size
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.NewUniform(color.Black))
	c.SetHinting(font.HintingFull)

	// Start from top with margin
	margin := 50
	lineHeight := int(c.PointToFixed(14) >> 6) // 1.4x line spacing
	pt := freetype.Pt(margin, margin+lineHeight)

	// Draw each line of text
	for _, text := range texts {
		_, err := c.DrawString(text, pt)
		if err != nil {
			return fmt.Errorf("failed to draw text: %v", err)
		}
		// Move to next line
		pt.Y += c.PointToFixed(14)
	}

	// Save the image
	return w.SaveImage(img, filepath.Base(outputPath))
}

// Write creates an image with sample text and saves it
func (w *ImageWriter) Write(outputPath string) error {
	return w.WriteTexts(outputPath, "Sample Text")
}
