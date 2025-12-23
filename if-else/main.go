package main

import (
	"fmt"

	"example.com/banl/fileops"
)

const fileName = "dummy_database.txt"

func main() {
	var accountBalance float64 = fileops.ReadValueFromFile(fileName)
	var choice int = 0

	for choice != 4 {

		fmt.Println("Dummy consonle Project!")
		fmt.Println("What do you want to do?")

		presentOptions()

		fmt.Println("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Yout balance:", accountBalance)
		case 2:
			var moneyToDepoist float64 = 0
			fmt.Println("How many?")
			fmt.Scan(&moneyToDepoist)
			accountBalance = accountBalance + moneyToDepoist
			fmt.Println("Yout balance:", accountBalance)
		case 3:
			var moneyToWitdhdraw float64 = 0
			fmt.Println("How many?")
			fmt.Scan(&moneyToWitdhdraw)
			accountBalance = accountBalance - moneyToWitdhdraw
			fmt.Println("Yout balance:", accountBalance)
		}
	}
	fileops.SaveValueInFile(accountBalance, fileName)

}
