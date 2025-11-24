package controllers

import (
	"net/http"

	"example.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(context *gin.Context){
	users, err := models.GetAllUsers()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"message": "Couldn't fetch users, try again later!",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message":"Users fetched successfully",
		"data":users,
	})
}