package models

import (
	"errors"
	"fmt"

	"github.com/MichaelVenturi/go-practice-api/db"
	"github.com/MichaelVenturi/go-practice-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users(email, password) VALUES (?, ?)
	`
	hash, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	res, err := utils.ExecQuery(db.DB, query, u.Email, hash)
	if err != nil {
		return err
	}

	userId, err := res.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := `
	SELECT id, password FROM users WHERE email = ?
	`
	row := db.DB.QueryRow(query, u.Email)

	var pass string
	err := row.Scan(&u.ID, &pass)
	if err != nil {
		fmt.Println(err)
		return errors.New("credentials invalid")
	}

	isValid := utils.CheckPasswordHash(u.Password, pass)
	if !isValid {
		return errors.New("credentials invalid")
	}

	return nil
}
