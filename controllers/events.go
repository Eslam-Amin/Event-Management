package controllers

import (
	"net/http"
	"strconv"

	"example.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func CreateEvent(context *gin.Context) {

	var event models.Event
	userId := context.GetInt64("userId")
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request data.",
			"error":   err.Error(),
		})
		return
	}

	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't create event, try again later!",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event created successfully",
		"data":    event,
	})
}

func GetAllEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Couldn't fetch events, try again later!",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": events,
	})
}

func GetEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request param.",
			"error":   err.Error(),
		})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Couldn't fetch event!",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event fetched successfully",
		"data":    event,
	})
}

func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request param.",
			"error":   err.Error(),
		})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Couldn't fetch event!",
			"error":   err.Error(),
		})
		return
	}

	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized to update this event.",
		})
		return
	}

	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request data.",
			"error":   err.Error(),
		})
		return
	}

	err = event.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't update event, try again later!",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "event updated Successfully.",
		"data":    event,
	})
}

func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request param.",
			"error":   err.Error(),
		})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Couldn't fetch event!",
			"error":   err.Error(),
		})
		return
	}

	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized to delete this event.",
		})
		return
	}
	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't delete event, try again later!",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted Successfully.",
		"data":    event,
	})
}

func RegisterForEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request param.",
			"error":   err.Error(),
		})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Couldn't fetch event!",
			"error":   err.Error(),
		})
		return
	}
	err = event.RegisterUserForEvent(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't register for event, try again later!",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Registered for event successfully",
	})

}

func CancelEventRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request param.",
			"error":   err.Error(),
		})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Couldn't fetch event!",
			"error":   err.Error(),
		})
		return
	}
	err = event.CancelEventRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't cancel event registration, try again later!",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event registration cancelled successfully",
	})

}
