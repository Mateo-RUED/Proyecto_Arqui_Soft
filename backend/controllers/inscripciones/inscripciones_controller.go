package controller_inscripciones

import (
    "net/http"
    dto_inscripciones "backend/dtos/inscripciones"
    inscripcionesService "backend/services/inscripciones"
    service_users "backend/services/users"
    "github.com/gin-gonic/gin"
    "strings"
)

// InscribirUsuario maneja las solicitudes de inscripción de usuarios en cursos.
func InscribirUsuario(context *gin.Context) {
    var request dto_inscripciones.InscribirUsuarioRequest
    if err := context.BindJSON(&request); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Solicitud inválida"})
        return
    }

    if err := inscripcionesService.InscribirUsuario(request); err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "Usuario inscrito correctamente"})
}

func ListCoursesByUser(context *gin.Context) {
    authHeader := context.GetHeader("Authorization")
    if authHeader == "" {
        context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
        return
    }

    tokenString := strings.TrimPrefix(authHeader, "Bearer ")
    userID, err := service_users.ValidateTokenAndGetUserID(tokenString)
    if err != nil {
        context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    response, err := inscripcionesService.ListCoursesByUser(userID)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, response)
}
