package routes

import (
	"todolist/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(router *gin.Engine) {
	router.GET("/todos", controllers.GetTodos)
	router.POST("/todos", controllers.CreateTodo)
	router.PUT("/todos/:id", controllers.UpdateTodo)
	router.DELETE("/todos/:id", controllers.DeleteTodo)

	router.GET("/users", controllers.GetUsers)
	router.POST("/signup", controllers.CreateUser)
	router.POST("/login", controllers.Login)
	// router.PUT("/users/:id", controllers.UpdateUser)
	// router.DELETE("/users/:id", controllers.DeleteUser)
}
