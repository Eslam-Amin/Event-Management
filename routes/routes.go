package routes

import (
	"example.com/event-booking/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){
	books := server.Group("/events")
	{
		books.GET("/", controllers.GetAllEvents)
		books.POST("/", controllers.CreateEvent)
		books.GET("/:id", controllers.GetEventById)
		books.PUT("/:id", controllers.UpdateEvent)
		books.DELETE("/:id", controllers.DeleteEvent)
	}

	auth := server.Group("/auth")
	{
		auth.POST("/signup", controllers.Singup)
		auth.POST("/login", controllers.Login)
	}

	users := server.Group("/users")
	{
		users.GET("/", controllers.GetAllUsers)
	}
}