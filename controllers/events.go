package controllers

import (
	"net/http"
	"strconv"

	"example.com/event-booking/models"
	"example.com/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func CreateEvent(context *gin.Context){
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message":"Unauthorized",
		})
		return
	}

	err := utils.ValidateToken(token)
	if err != nil{
		context.JSON(http.StatusUnauthorized, gin.H{
			"message":"Unauthorized",
			"error": err.Error(),
		})
		return
	}

  var event models.Event

  err = context.ShouldBindJSON(&event)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{
            "message": "Couldn't parse request data.",
						"error": err.Error(),
					})
					return
				}
				
				event.UserID = 1
				
				err = event.Save()
				if err != nil {
					context.JSON(http.StatusInternalServerError, gin.H{
						"message":"Couldn't create event, try again later!",
						"error": err.Error(),
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
						"error": err.Error(),
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
						"error": err.Error(),
					})
					return
				}
				event, err := models.GetEventById(eventId)
				if err != nil{
					context.JSON(http.StatusNotFound, gin.H{
						"message": "Couldn't fetch event!",
						"error": err.Error(),
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
						"error": err.Error(),
					})
					return
				}
				
				event, err := models.GetEventById(eventId)
				if err != nil{
					context.JSON(http.StatusNotFound, gin.H{
						"message": "Couldn't fetch event!",
						"error": err.Error(),
					})
					return
				}
				
				err = context.ShouldBindJSON(&event)
				if err != nil{
					context.JSON(http.StatusBadRequest, gin.H{
						"message": "Couldn't parse request data.",
						"error": err.Error(),
					})
					return
				}
				
				err = event.Update()
				if err != nil{
					context.JSON(http.StatusInternalServerError, gin.H{
						"message": "Couldn't update event, try again later!",
						"error": err.Error(),
					})
					return
				}
				context.JSON(http.StatusOK, gin.H{
					"message": "event updated Successfully.",
					"data": event,
				})
			}
			
			func DeleteEvent(context *gin.Context){
				eventId, err := strconv.ParseInt(context.Param("id"),10,64)
				if err != nil{
					context.JSON(http.StatusBadRequest, gin.H{
						"message": "Couldn't parse request param.",
						"error": err.Error(),
					})
					return
				}
				
				event, err := models.GetEventById(eventId)
				if err != nil{
					context.JSON(http.StatusNotFound, gin.H{
						"message": "Couldn't fetch event!",
						"error": err.Error(),
					})
					return
				}
				
				err = event.Delete()
				
				if err != nil{
					context.JSON(http.StatusInternalServerError, gin.H{
						"message": "Couldn't delete event, try again later!",
						"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted Successfully.",
		"data": event,
	})
}



