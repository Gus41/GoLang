package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	User

	email    string
	password string
}

func NewAdmin(email, pass string, u User) Admin {
	return Admin{
		password: pass,
		email:    email,
		User:     u,
	}
}

func (user *User) OutputUserDetails() {
	fmt.Println("=======================")
	fmt.Println(user.firstName, user.lastName, user.birthDate, user.createdAt)

}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func New(firstName, lastName, birthDate string) (*User, error) {

	if firstName != "" && lastName != "" && birthDate != "" {
		return &User{
			firstName: firstName,
			lastName:  lastName,
			birthDate: birthDate,
			createdAt: time.Now(),
		}, nil
	}

	return nil, errors.New("Invalid data")

}
