package controller_users

import (
	dto_users "backend/dtos/users"
	usersService "backend/services/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var loginRequest dto_users.LoginRequest
	if err := context.BindJSON(&loginRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	response, err := usersService.Login(loginRequest)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, response)
}

func CreateUser(context *gin.Context) {
	var createUserRequest dto_users.CreateUserRequest
	if err := context.BindJSON(&createUserRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := usersService.CreateUser(createUserRequest); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
