package controller_inscripciones

import (
    "net/http"
    dto_inscripciones "backend/dtos/inscripciones"
    inscripcionesService "backend/services/inscripciones"
    "github.com/gin-gonic/gin"
    "strconv"
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
    // Obtener el parametro usuarioID de la URL
    usuarioIDParam := context.Param("usuarioID")

    // Convertir el parametro usuarioID a uint
    usuarioID, err := strconv.ParseUint(usuarioIDParam, 10, 32)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // Llamar al servicio con el usuarioID convertido
    response, err := inscripcionesService.ListCoursesByUser(uint(usuarioID))
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, response)
}