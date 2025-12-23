package main

import (
	"fmt"
)

func getAdultYears(age *int) {
	*age = *age - 18

}

func main() {
	age := 32
	var agePointer *int = &age

	fmt.Println("Age value: ", *agePointer)
	fmt.Println("Age address: ", agePointer)
	getAdultYears(agePointer)
	fmt.Println("Age Adullt years: ", age)
	fmt.Println("Age value: ", age)

}
