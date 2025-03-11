package models

import (
	"errors"
	"log"
	"todolist/config"
)

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func CreateUser(user User) {
	query := "INSERT INTO users (name, email, password, role) VALUES (?, ?, ?, ?)"
	_, err := config.DB.Exec(query, user.Name, user.Email, user.Password, user.Role)
	if err != nil {
		log.Println("Failed to insert user: ", err)
	}
}

func GetUsers() []User {
	rows, err := config.DB.Query("SELECT id, name, email, password, role FROM users")
	if err != nil {
		log.Println("Failed to query users: ", err)
		return nil
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role); err != nil {
			log.Println("Failed to scan user: ", err)
			return nil
		}
		users = append(users, user)
	}

	return users
}

func Verify(user User) (bool, error) {
	query := "SELECT password FROM users WHERE email = ?"
	row := config.DB.QueryRow(query, user.Email)
	var password string
	err := row.Scan(&password)
	if err != nil {
		return false, errors.New("User not found")
	}
	if password != user.Password {
		return false, errors.New("Password is incorrect")
	}
	return true, nil
}
