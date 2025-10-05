package main

import "fmt"

func main() {
	// %T -> Type
	//%v -> value

	fmt.Printf("Type: %T - value: %v \n", true, true)
	fmt.Printf("Type: %T - value: %v \n", "test", "Test")
	fmt.Printf("Type: %T - value: %v \n", 1, 1)
	fmt.Printf("Type: %T - value: %v \n", "1", "1")
	fmt.Printf("Type: %T - value: %v \n", 1.123, 1.123)
}
