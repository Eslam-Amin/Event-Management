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

	server.POST("/events", func (context *gin.Context){
		var event models.Event
		err := context.ShouldBindJSON(&event)
		fmt.Println("err: ",err)
		if err != nil {	
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
		return
		}

		event.ID = 1
		event.UserID = 1
		event.Save()
		context.JSON(http.StatusCreated, gin.H{
			"message":"Event created successfully",
			"data": event,
		})



	})

	server.GET("/events", func (context *gin.Context){
		context.JSON(http.StatusOK, gin.H{
			"data": models.GetAllEvents(),
		})
	})

	server.Run(":8080")
}