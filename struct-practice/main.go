package main

import (
	"bufio"
	"fmt"
	"noteproject/do"
	"noteproject/note"
	"os"
	"strings"
)

type saver interface {
	Save() error
}
type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func saveData(data outputtable) {
	err := data.Save()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data saved")
}

func main() {
	title, content := getNoteData()
	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todoText := getUserInput("Todo: ")
	todo, err := do.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(note)
	outputData(todo)

}

func outputData(data outputtable) {
	data.Display()
	saveData(data)
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
