package routes

import (
	"net/http"

	"example.com/web/models"
	"example.com/web/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	user := models.User{}
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't get user"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't save user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Save user successfully!", "user: ": user})

}

func login(context *gin.Context) {
	user := models.User{}
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't get user"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		panic(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})
}

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't get users"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"users": users})
}
