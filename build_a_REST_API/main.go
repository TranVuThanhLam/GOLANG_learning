package main

import (
	"example.com/web/db"
	"example.com/web/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRouter(server)

	server.Run(":8234")

}
