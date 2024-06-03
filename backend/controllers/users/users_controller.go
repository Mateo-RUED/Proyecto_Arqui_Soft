package controller_users

import (
    dto_users"backend/dtos/users"
    "net/http"
    usersDomain "backend/domain/users"
    usersService "backend/services/users"
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

// CreateUser handles user creation requests.
func CreateUser(context *gin.Context) {
    var request dto_users.CreateUserRequest
    if err := context.BindJSON(&request); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Solicitud inválida"})
        return
    }

    user := usersDomain.User{
        Username: request.Username,
        Password: request.Password,
        Tipo:     request.Tipo,
        Email:    request.Email,
    }

    if err := usersService.CreateUser(user); err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusCreated, gin.H{"message": "Usuario creado correctamente"})
}

// GenerateToken handles the token generation.
func GenerateToken(c *gin.Context) {
    var request dto_users.TokenRequest
    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    tokenString, err := usersService.GenerateJWT(request.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}





