package main

import (
	"todolist/config"
	"todolist/models"
	"todolist/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Todo{})

	r := gin.Default()

	// Middleware CORS để React có thể kết nối
	r.Use(cors.Default())

	routes.TodoRoutes(r)

	r.Run(":8080")
}
