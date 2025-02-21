package routes

import "github.com/gin-gonic/gin"

func RegisterRouter(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/event_:id", getEventById)

	server.POST("/event", createEvent)
	server.POST("/signup", signup)

	server.PUT("/event_:id", updateEvent)
	server.DELETE("/event_:id", deleteEvent)
}
