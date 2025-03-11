package routes

import (
	"example.com/web/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/event_:id", getEventById)
	server.GET("/users", getUsers)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/event", middlewares.Authenticate, createEvent)
	authenticated.PUT("/event_:id", updateEvent)
	authenticated.DELETE("/event_:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
