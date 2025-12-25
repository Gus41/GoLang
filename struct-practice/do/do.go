package do

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content string `json:"content"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Content)
}

func (todo Todo) Save() error {

	jsonContent, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile("todo.json", jsonContent, 0644)
}

func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("Invalid input")
	}

	return Todo{
		Content: content,
	}, nil
}
