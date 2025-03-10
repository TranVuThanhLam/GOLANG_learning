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
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at"`
}

func (note Note) Display() {
	fmt.Printf("Your note Title %v has the following Content: \n\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_") + ".json"
	fileName = strings.ToLower(fileName)
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" {
		return Note{}, errors.New("Invalid input Title.")
	}
	if content == "" {
		return Note{}, errors.New("Invalid input Content.")
	}
	return Note{
		Title:    title,
		Content:  content,
		CreateAt: time.Now(),
	}, nil
}
