package controller_courses

import (
    "net/http"
    dto_courses"backend/dtos/courses"
    coursesService"backend/services/courses"
    "github.com/gin-gonic/gin"
    "strconv"
)

// CreateCourse maneja las solicitudes de creación de cursos.
func CreateCourse(context *gin.Context) {
    var request dto_courses.CreateCourseRequest
    if err := context.BindJSON(&request); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Solicitud inválida"})
        return
    }

    if err := coursesService.CreateCourse(request); err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusCreated, gin.H{"message": "Curso creado exitosamente"})
}

func DeleteCourseName(context *gin.Context) {
    var request dto_courses.DeleteCourseRequest
    if err := context.BindJSON(&request); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Solicitud inválida"})
        return
    }

    if err := coursesService.DeleteCurso(request.Name); err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "Curso eliminado correctamente"})
}


func EditCourse(context *gin.Context) {
    var editRequest dto_courses.EditCourseRequest
    if err := context.BindJSON(&editRequest); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Solicitud inválida"})
        return
    }

    // Llamar al servicio para editar el curso
    if err := coursesService.EditCourse(editRequest); err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "Curso editado exitosamente"})
}

func GetCourseByID(context *gin.Context) {
    var request dto_courses.GetCourseByIDRequest
    if err := context.ShouldBindUri(&request); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    response, err := coursesService.GetCourseByID(request.ID)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, response)
}

func GetCourseByIDParam(context *gin.Context) {
    idParam := context.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    response, err := coursesService.GetCourseByID(uint(id))
    if err != nil {
        context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, response)
}

func DeleteCourseByIDParam(context *gin.Context) {
    idParam := context.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    err = coursesService.DeleteCourseByID(uint(id))
    if err != nil {
        context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}

func GetCoursesByCategory(context *gin.Context) {
    category := context.Query("category")
    if category == "" {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Category is required"})
        return
    }

    response, err := coursesService.GetCoursesByCategory(category)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"courses": response})
} 

func GetAllCourses(context *gin.Context) {
    response, err := coursesService.GetAllCourses()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusOK, gin.H{"courses": response})
}
