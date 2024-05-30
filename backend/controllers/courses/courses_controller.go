/* package controllers

import (
        "net/http"
        coursesDomain "backend/domain/courses"
        coursesService "backend/services/courses"
        "github.com/gin-gonic/gin"
    )

type CourseController struct {
    Service *services.CourseService
}

func NewCourseController(service *services.CourseService) *CourseController {
    return &CourseController{
        Service: service,
    }
}

func (cc *CourseController) CreateCourse(c *gin.Context) {
    var courseDTO dtos.CreateCourseDTO
    if err := c.ShouldBindJSON(&courseDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    course, err := cc.Service.CreateCourse(courseDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, course)
}

func (cc *CourseController) UpdateCourse(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var courseDTO dtos.UpdateCourseDTO
    if err := c.ShouldBindJSON(&courseDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    course, err := cc.Service.UpdateCourse(uint(id), courseDTO)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, course)
}

func (cc *CourseController) DeleteCourse(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := cc.Service.DeleteCourse(uint(id)); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
} */


// controllers/course_controller.go
package controllers

import (
    "Proyecto_Arqui_Soft/backend/dtos"
    "Proyecto_Arqui_Soft/backend/services"
    "Proyecto_Arqui_Soft/backend/utils/errors"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type CourseController struct {
    Service services.CourseServiceInterface
}

func NewCourseController(service services.CourseServiceInterface) *CourseController {
    return &CourseController{
        Service: service,
    }
}

func (cc *CourseController) CreateCourse(c *gin.Context) {
    var courseDTO dtos.CreateCourseDTO
    if err := c.ShouldBindJSON(&courseDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    course, apiErr := cc.Service.CreateCourse(courseDTO)
    if apiErr != nil {
        c.JSON(apiErr.Status(), apiErr)
        return
    }

    c.JSON(http.StatusCreated, course)
}

func (cc *CourseController) UpdateCourse(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error":        "invalid course ID"})
        return
    }

    var courseDTO dtos.UpdateCourseDTO
    if err := c.ShouldBindJSON(&courseDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    course, apiErr := cc.Service.UpdateCourse(uint(id), courseDTO)
    if apiErr != nil {
        c.JSON(apiErr.Status(), apiErr)
        return
    }

    c.JSON(http.StatusOK, course)
}

func (cc *CourseController) DeleteCourse(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
        return
    }

    apiErr := cc.Service.DeleteCourse(uint(id))
    if apiErr != nil {
        c.JSON(apiErr.Status(), apiErr)
        return
    }

    c.Status(http.StatusNoContent)
}

func (cc *CourseController) GetCourseById(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
        return
    }

    course, apiErr := cc.Service.GetCourseById(uint(id))
    if apiErr != nil {
        c.JSON(apiErr.Status(), apiErr)
        return
    }

    c.JSON(http.StatusOK, course)
}

func (cc *CourseController) GetCourses(c *gin.Context) {
    courses, apiErr := cc.Service.GetCourses()
    if apiErr != nil {
        c.JSON(apiErr.Status(), apiErr)
        return
    }

    c.JSON(http.StatusOK, courses)
}



