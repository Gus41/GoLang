package main

import (
	"fmt"
	"structsProject/user"
)

type CustomString string

func (s CustomString) Log() {
	fmt.Println(s)
}

func main() {

	firstName := getUserData("First name ")
	lastName := getUserData("Last name ")
	birthDate := getUserData("Birth Date ")

	appUser, err := user.New(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	var test CustomString = "test"

	test.Log()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)

	return value
}
