package routes

import (
	"example.com/event-booking/controllers"
	"example.com/event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	books := server.Group("/events")
	{
		books.GET("/", controllers.GetAllEvents)
		books.GET("/:id", controllers.GetEventById)
		protectedBooks := books.Group("/")
		protectedBooks.Use(middlewares.Authticate)
		protectedBooks.Use(middlewares.Authorize)
		protectedBooks.POST("/", controllers.CreateEvent)
		protectedBooks.PUT("/:id", controllers.UpdateEvent)
		protectedBooks.DELETE("/:id", controllers.DeleteEvent)
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
