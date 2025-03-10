package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/web/models"
	"github.com/gin-gonic/gin"
)

func createEvent(context *gin.Context) {
	event := models.Event{}
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	userId := context.GetInt64("userId")

	event.UserID = int(userId)

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save events"})
		// panic(err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event: ": fmt.Sprint(event)})
}

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to convert id to integer"})
		return
	}

	event, err := models.GetEventById(eventId)
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

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to convert id to integer"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Event doesn't exists"})
		return
	}

	userId := context.GetInt64("userId")

	if event.UserID != int(userId) {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Only event's creater can update"})
	}

	updatedEvent := models.Event{}
	// // get data from POST or PUT method
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not get data."})
		return
	}

	updatedEvent.ID = eventId

	err = models.UpdateEvents(eventId, &updatedEvent)
	if err != nil {
		context.JSON(http.StatusAlreadyReported, gin.H{"message": "Failed to update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "update successful"})
}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message:": "Couldn't get id"})
		return
	}

	event, err := models.GetEventById(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "event doesn't exists!"})
	}

	userId := context.GetInt64("userId")

	if event.UserID != int(userId) {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Only event's creater can delete"})
	}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message:": "Couldn't delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Delete successfully!"})
}
