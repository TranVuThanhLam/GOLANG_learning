package main

import (
	"todolist/config"
	"todolist/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	config.InitDB()

	// khoi tao routera
	router := gin.Default()

	// Middleware CORS để React có thể kết nối
	router.Use(cors.Default())

	routes.TodoRoutes(router)

	router.Run(":8080")
}
