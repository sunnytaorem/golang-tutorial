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
	Title     string    `json:"title"` // struct tags are meta data that specify how the fields should be encoded/decoded when converting to/from JSON. In this case, the Title field will be represented as "title" in the JSON data.
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// Since it doesn't edit the Note struct, we can use a value receiver instead of a pointer receiver.
// This means that the method will receive a copy of the Note struct, and any changes made to it will not affect the original Note instance.
func (n Note) Display() {
	fmt.Printf("Your note titled %v has the following content: \n\n%v\n", n.Title, n.Content)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	// json.Marshal extracts data from the Note struct that are exported (public) and converts it into a JSON format.
	// It can also read the meta data specified in the struct tags to determine how to represent the fields in the JSON output.
	// It returns the JSON data as a byte slice and an error if any occurs during the marshaling process.
	jsonData, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}

// Constructor function to create a new Note instance with validation.
func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("Title and content cannot be empty")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
