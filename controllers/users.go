package controllers

import (
	"net/http"
	"strconv"

	"example.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Couldn't fetch users, try again later!",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Users fetched successfully",
		"data":    users,
	})
}

func GetUserById(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid user ID",
		})
		return
	}
	user, err := models.GetUserById(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Couldn't fetch user, try again later!",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"data":    user,
	})
}

func GetUsersRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	registrations, err := models.GetRegistrationsByUserId(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Couldn't fetch registrations, try again later!",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Registrations fetched successfully",
		"data":    registrations,
	})
}
