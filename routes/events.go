package routes

import (
	"net/http"
	"strconv"

	"github.com/artur-sres/golang-event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"events": events})
}

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event ID."})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"event": event})
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event ID."})
		return
	}

	_, err = models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	updatedEvent.ID = eventID

	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event ID."})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully", "event": event})
}
