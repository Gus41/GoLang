package main

import (
	"bufio"
	"fmt"
	"noteproject/do"
	"noteproject/note"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()
	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()

	err = note.Save()
	if err != nil {
		fmt.Println(err)
	}

	todoText := getUserInput("Todo: ")
	todo, err := do.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = todo.Save()

	if err != nil {
		fmt.Print(err)
	}
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	rd := bufio.NewReader(os.Stdin)
	value, err := rd.ReadString('\n')

	if err != nil {
		return ""
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

	return value
}
