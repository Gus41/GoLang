package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}

	dbNumbers := transformNumbers(&numbers, multiplyBy(2))
	trNumbers := transformNumbers(&numbers, func(n int) int {
		return n * 3
	})

	fmt.Println("Original:", numbers)
	fmt.Println("Dobrados:", dbNumbers)
	fmt.Println("Triplicados:", trNumbers)

}

type transformFunction func(int) int

func multiplyBy(num int) transformFunction {
	return func(value int) int {
		return num * value
	}
}

func transformNumbers(numbers *[]int, transform transformFunction) []int {
	dbNumbers := make([]int, len(*numbers))
	for i, val := range *numbers {
		dbNumbers[i] = transform(val)
	}

	return dbNumbers
}
