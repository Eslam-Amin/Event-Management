package main

import (
	"fmt"
	"net/http"

	"example.com/event-booking/db"
	"example.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/ping", func (context *gin.Context){
	
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.POST("/events", createEvent)
	server.GET("/events", getAllEvents)

	server.Run(":8080")
}

func createEvent(context *gin.Context){
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



func getAllEvents(context *gin.Context) {
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