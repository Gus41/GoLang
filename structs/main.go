package main

import (
	"fmt"
	"structsProject/user"
)

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

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)

	return value
}
