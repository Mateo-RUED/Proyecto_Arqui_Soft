package controller_courses

import (
	dto_courses "backend/dtos/courses"
	services_courses "backend/services/courses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddComment(context *gin.Context) {
	var request dto_courses.AddCommentRequest

	// Validar la solicitud
	if err := context.BindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Solicitud inválida: verifica los datos enviados."})
		return
	}

	// Intentar agregar el comentario
	if err := services_courses.AddComment(request); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudo agregar el comentario: " + err.Error()})
		return
	}

	// Respuesta exitosa
	context.JSON(http.StatusOK, gin.H{"message": "Comentario agregado exitosamente."})
}

func GetCommentsByCourse(context *gin.Context) {
	// Convertir el ID del curso a entero
	courseID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "ID de curso inválido."})
		return
	}

	// Obtener comentarios del curso
	comments, err := services_courses.GetCommentsByCourse(uint(courseID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No se pudieron obtener los comentarios: " + err.Error()})
		return
	}

	// Respuesta exitosa con los comentarios
	context.JSON(http.StatusOK, gin.H{"comments": comments})
}
