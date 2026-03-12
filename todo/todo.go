package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (t Todo) Display() {
	fmt.Println(t.Text)
}

func (t Todo) Save() error {
	fileName := "todo.json"

	// json.Marshal extracts data from the Todo struct that are exported (public) and converts it into a JSON format.
	// It can also read the meta data specified in the struct tags to determine how to represent the fields in the JSON output.
	// It returns the JSON data as a byte slice and an error if any occurs during the marshaling process.
	jsonData, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}

// Constructor function to create a new Todo instance with validation.
func New(content string) (*Todo, error) {
	if content == "" {
		return &Todo{}, errors.New("Text cannot be empty")
	}
	return &Todo{
		Text: content,
	}, nil
}
