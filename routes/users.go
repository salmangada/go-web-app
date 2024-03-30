package routes

import (
	"net/http"

	"com.sal/main/models"
	"com.sal/main/utils"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse data"})
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to save User"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func login(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse data"})
		return
	}

	err = user.ValidateCreds()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authorize user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authorize user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login SUccesssful", "toke,": token})
}
