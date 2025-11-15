package main

import (
	"fmt"
	"os"
	"strconv"
)

const fileName = "dummy_database.txt"

func saveValueInFile(value float64) {
	var valuteText = fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valuteText), 0644)
}
func readValueFromFile() float64 {
	data, err := os.ReadFile(fileName)
	if err != nil {
		saveValueInFile(0)
		return 0.0
	}

	valueText := string(data)
	valueFloat, err := strconv.ParseFloat(valueText, 64)

	//nil is basicly null (???)
	if err != nil {
		return 0.0
	}

	return valueFloat
}

func main() {
	var accountBalance float64 = readValueFromFile()
	var choice int = 0

	for choice != 4 {

		fmt.Println("Dummy consonle Project!")
		fmt.Println("What do you want to do?")

		fmt.Println("1 - Check balance")
		fmt.Println("2 - Deposit Money")
		fmt.Println("3 - Withdraw money")
		fmt.Println("4 - Exit")

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
	saveValueInFile(accountBalance)

}
