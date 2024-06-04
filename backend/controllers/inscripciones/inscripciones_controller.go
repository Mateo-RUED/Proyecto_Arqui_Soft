package controller_inscripciones

import (
    "net/http"
    "backend/dtos/inscripciones"
    inscripcionesService "backend/services/inscripciones"
    "github.com/gin-gonic/gin"
)

func ListCoursesByUser(context *gin.Context) {
    var request dto_inscripciones.ListCoursesByUserRequest
    if err := context.BindJSON(&request); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    response, err := inscripcionesService.ListCoursesByUser(request.UsuarioID)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, response)
}
