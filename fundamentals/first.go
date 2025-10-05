package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world!")

	var output string
	output = strings.Split("Teste - teste", "-")[0]
	output = strings.TrimSpace(output)
	fmt.Println(output)
}
