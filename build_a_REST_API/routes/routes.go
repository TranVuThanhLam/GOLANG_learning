package routes

import "github.com/gin-gonic/gin"

func RegisterRouter(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/event_:id", getEventById)
	server.GET("/users", getUsers)

	server.POST("/event", createEvent)
	

	server.PUT("/event_:id", updateEvent)
	server.DELETE("/event_:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
