package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("sqlite3", "todolist.db")
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping database")
	}
	log.Println("Database connected!")
}

func InitDB() {
	// Create the todos table if it doesn't exist
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS todos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			status BOOLEAN NOT NULL
		)`)
	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}

	// Create the todos table if it doesn't exist
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		role TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}
}
