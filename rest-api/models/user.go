package models

import (
	"api/db"
	"api/utils"
)

type User struct {
	Id       int
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users(email,password) VALUES (?,?)	
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPass)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.Id = int(userId)

	return err

}
