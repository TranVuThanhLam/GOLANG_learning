package routes

import "github.com/gin-gonic/gin"

func RegisterRouter(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/event_by_id_:id", getEventById)

	server.POST("/events", createEvents)
	server.PUT("/event_update/:id", updateEvents)

}
