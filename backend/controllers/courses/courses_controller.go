package controllers

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
    var courseDTO dto.CreateCourseDTO
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
    var courseDTO dto.UpdateCourseDTO
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
}


