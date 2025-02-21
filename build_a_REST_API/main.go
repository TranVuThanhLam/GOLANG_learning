package main

import (
	"fmt"
	"net/http"

	"example.com/web/db"
	"example.com/web/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/event_by_id_:id", getEventById)

	server.POST("/events", createEvents)

	server.Run()

}

func getEventById(context *gin.Context) {
	id := context.Param("id")
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get all event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get all event"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	event := models.Event{}
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		// context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save events"})
		// context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		panic(err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event: ": fmt.Sprint(event)})
}
