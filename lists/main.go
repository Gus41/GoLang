package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {

	forFunction()

	fmt.Println("Hello world!")

	prices := [4]float64{11.3, 33.1, 14.2, 99.9}
	fmt.Println(prices)

	var productNames []string
	productNames = append(productNames, "teste")
	fmt.Print(productNames)

	//slecedPrices := prices[1:3] //first value will allaways be included but no the second (0,1,2)
	//fmt.Print(slecedPrices)

}

func maps() {

	//websites := []string{"https://google.com", "https://aws.com"}

	websitesMap := map[string]string{
		"google":             "https:google.com",
		"amazon web service": "https://aws.com",
	}

	fmt.Print(websitesMap)
	fmt.Println(websitesMap["google"])

	websitesMap["new"] = "new value"
	delete(websitesMap, "new")

}

func makeFunction() {
	userNames := make([]string, 1, 5)
	userNames[0] = "initialuser"
	userNames = append(userNames, "user")

	fmt.Println(userNames)

	dict := map[string]int{}
	dict["0"] = 0

	type floatMap map[string]float64

	fm := floatMap{}
	fm["test"] = 1.2
}

func forFunction() {
	// Criamos um slice com tamanho 0, mas capacidade para 10
	items := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		items = append(items, i) // Não haverá novas alocações de memória aqui!
	}
	for i, value := range items {
		fmt.Println("Index:", i)
		fmt.Println("value:", value)
	}
}
