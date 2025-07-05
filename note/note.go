package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note *Note) Display() {
	fmt.Printf("\n###\nTitle: \"%s\"\nContent: \"%s\"\n###\n", note.Title, note.Content)
}

func (note *Note) Save() error {
	jsonData, err := json.Marshal(note)

	if err != nil {
		return err
	}

	err = os.WriteFile(
		strings.ToLower(strings.ReplaceAll(note.Title, " ", "_"))+".json",
		jsonData,
		644,
	)

	if err != nil {
		return err
	}

	return nil
}

func New(title string, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("invalid note data")
	}

	return &Note{Title: title, Content: content, CreatedAt: time.Now()}, nil
}
