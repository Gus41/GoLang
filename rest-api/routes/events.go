package routes

import (
	"api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"Message": "Could not fetch data",
			},
		)

		return
	}
	context.JSON(
		http.StatusOK,
		gin.H{
			"events": events,
		},
	)
}

func updateEvent(context *gin.Context) {
	id := context.Param("id")
	eventId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Invalid Param",
			},
		)
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"Message": "Could not fetch event",
			},
		)

		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Could not parse request data",
			},
		)

		return
	}

	updatedEvent.Id = int(eventId)

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"Message": "Could not update data",
			},
		)

		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"event": updatedEvent,
		},
	)

}

func getEventById(context *gin.Context) {
	id := context.Param("id")
	eventId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Invalid Param",
			},
		)
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"Message": "Could not fetch event",
			},
		)

		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"event": event,
		},
	)

}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"Message": "Could not parse request data",
			},
		)

		return
	}

	event.Id = 1
	event.UserId = 1

	err = event.Save()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"Message": "Could not create event",
			},
		)

		return
	}
	context.JSON(
		http.StatusCreated,
		gin.H{
			"Message": "Event Created",
			"Event":   event,
		},
	)

}
