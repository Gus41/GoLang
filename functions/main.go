package main

import (
	"fmt"
)

func main() {
	outputText("Hello world")
	value := fmt.Sprint(calculate(1, 1))
	outputText(value)
}

func outputText(text string) {
	fmt.Println(text)
}
func calculate(a int, b int) int {
	return a + b
}
