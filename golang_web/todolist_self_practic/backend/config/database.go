package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("todolist.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	log.Println("Database connected!")
}
