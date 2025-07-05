package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const fileName = "todo.json"

type Todo struct {
	Text string `json:"title"`
}

func (todo *Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo *Todo) Save() error {
	jsonData, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	if err = os.WriteFile(fileName, jsonData, 644); err != nil {
		return err
	}

	return nil
}

func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("invalid todo data")
	}

	return &Todo{Text: text}, nil
}
