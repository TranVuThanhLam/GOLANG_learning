package controllers

import (
	"todolist/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	user := models.User{}
	context.ShouldBindJSON(&user)
	models.CreateUser(user)
	context.JSON(200, gin.H{"message": "User created successfully"})
}

func GetUsers(context *gin.Context) {
	users := models.GetUsers()
	context.JSON(200, users)
}

func Login(context *gin.Context) {
	user := models.User{}
	context.ShouldBindJSON(&user)
	ok, err := models.Verify(user)
	if err != nil {
		context.JSON(401, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		context.JSON(401, gin.H{"error": "User not verified"})
		return
	}
	context.JSON(200, gin.H{"message": "User logged in successfully"})
}
