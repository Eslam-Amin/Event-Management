package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func CreateEvent(context *gin.Context){
  var event models.Event

    err := context.ShouldBindJSON(&event)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{
            "message": "Couldn't parse request data.",
        })
        return
    }

    event.UserID = 1

    err = event.Save()
    if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusInternalServerError, gin.H{
            "message":"Couldn't create event, try again later!",
        })
        return
    }

    context.JSON(http.StatusOK, gin.H{
        "message": "Event created successfully",
        "data": event,
    })
}

func GetAllEvents(context *gin.Context) {
	events, err :=  models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't fetch events, try again later!",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": events,
	})
}

func GetEventById(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10 , 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request param.",
		})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't fetch event, try again later!",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message":"Event fetched successfully",
		"data":event,
	})
}

func UpdateEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request param.",
		})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't fetch event, try again later!",
		})
		return
	}

	err = context.ShouldBindJSON(&event)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't parse request data.",
		})
		return
	}

	err = event.Update()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't update event, try again later!",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "event updated Successfully.",
		"data": event,
	})
}



