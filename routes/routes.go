package routes

import (
	"example.com/event-booking/controllers"
	"example.com/event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	events := server.Group("/events")
	{
		events.GET("/", controllers.GetAllEvents)
		events.GET("/:id", controllers.GetEventById)
		protectedEvents := events.Group("/")
		protectedEvents.Use(middlewares.Authticate)
		protectedEvents.Use(middlewares.Authorize)
		protectedEvents.POST("/", controllers.CreateEvent)
		protectedEvents.PUT("/:id", controllers.UpdateEvent)
		protectedEvents.DELETE("/:id", controllers.DeleteEvent)
		protectedEvents.POST("/:id/register", controllers.RegisterForEvent)
		protectedEvents.DELETE("/:id/register", controllers.CancelEventRegistration)

	}

	auth := server.Group("/auth")
	{
		auth.POST("/signup", controllers.Singup)
		auth.POST("/login", controllers.Login)
	}

	users := server.Group("/users")
	{
		users.GET("/", controllers.GetAllUsers)
		users.GET("/:id", controllers.GetUserById)
	}
}
