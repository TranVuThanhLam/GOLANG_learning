package models

import (
	"errors"

	"example.com/web/db"
	"example.com/web/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users (email, password)
	VALUES(?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(u.Email, hashPassword)

	if err != nil {
		panic(err)
	}
	return nil
}

// Credential : thong tin da xac thuc
func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrivedPassword string
	err := row.Scan(&u.Id, &retrivedPassword)

	if err != nil {
		return errors.New("Gmail sai!")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrivedPassword)

	if !passwordIsValid {
		return errors.New("Passs sai!")
	}

	return nil
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, errors.New("Can't reading data!")
	}

	users := []User{}
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Email, &user.Password)
		users = append(users, user)
	}
	return users, nil
}
