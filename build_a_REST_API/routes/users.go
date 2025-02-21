package routes

import (
	"net/http"

	"example.com/web/models"
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
