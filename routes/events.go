package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MichaelVenturi/go-practice-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	fmt.Println("Get request made")
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {

	userId := context.GetInt64("userId") // added to context object thru auth middleware

	var event models.Event
	err := context.ShouldBindJSON(&event) // automatically puts request body into given variable if it shares this variable's structure
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	event.UserID = userId

	err = event.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId) // confirm event exists
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	// check if user owns this event
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized for this event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event updated!", "updated_event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId) // confirm event exists
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	// check if user owns this event
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized for this event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
