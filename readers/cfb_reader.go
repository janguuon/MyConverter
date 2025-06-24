package readers

import (
	"fmt"
	"io"
	"os"

	"github.com/richardlehane/mscfb"
)

// CFBReader implements FileReader for CFB format
type CFBReader struct {
	FilePath string
	Entries  []CFBEntry
}

// CFBEntry represents an entry in the CFB file
type CFBEntry struct {
	Name     string
	Size     int64
	Content  []byte
	Children []*CFBEntry
}

// Read implements CFB file reading
func (r *CFBReader) Read(filePath string) error {
	r.FilePath = filePath
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Parse the CFB file
	reader, err := mscfb.New(file)
	if err != nil {
		return fmt.Errorf("failed to parse CFB file: %v", err)
	}

	// Read all entries
	for entry, err := reader.Next(); err != io.EOF; entry, err = reader.Next() {
		if err != nil {
			return fmt.Errorf("error reading entry: %v", err)
		}

		content := make([]byte, entry.Size)
		_, err = reader.Read(content)
		if err != nil {
			return fmt.Errorf("error reading entry content: %v", err)
		}

		r.Entries = append(r.Entries, CFBEntry{
			Name:    entry.Name,
			Size:    entry.Size,
			Content: content,
		})
	}

	return nil
}

// GetEntries returns all entries in the CFB file
func (r *CFBReader) GetEntries() []CFBEntry {
	return r.Entries
}
