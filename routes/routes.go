package routes

import (
	"example.com/event-booking/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){
	server.GET("/events", controllers.GetAllEvents)
	server.GET("/events/:id", controllers.GetEventById)
	server.POST("/events", controllers.CreateEvent)
	server.PUT("/events/:id", controllers.UpdateEvent)
}