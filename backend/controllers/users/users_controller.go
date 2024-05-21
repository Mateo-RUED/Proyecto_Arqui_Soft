package users

import (
    "net/http"
    usersDomain "backend/domain/users"
    usersService "backend/services/users"
    "github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
    var loginRequest usersDomain.LoginRequest
    if err := context.BindJSON(&loginRequest); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid requesta"})
        return
    }
    response := usersService.Login(loginRequest)
    context.JSON(http.StatusOK, response)
}

// GenerateToken handles the token generation.
func GenerateToken(c *gin.Context) {
    var request usersDomain.TokenRequest
    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid requeste"})
        return
    }

    tokenString, err := usersService.GenerateJWT(request.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}


