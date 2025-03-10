package routes

import (
	"todolist/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.Engine) {
	r.GET("/todos", controllers.GetTodos)
	r.POST("/todos", controllers.CreateTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)
}
