package courses

import (
    "net/http"
    // dto_courses"backend/dtos/courses"
    CoursesService"backend/services/courses"
    "github.com/gin-gonic/gin"
)

type CourseController struct{}

func NewCourseController() *CourseController {
    return &CourseController{}
}

func (cc *CourseController) CreateCourse(c *gin.Context) {
    var courseDTO CoursesService.CreateCourseDTO
    if err := c.ShouldBindJSON(&courseDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    course, err := CoursesService.CreateCourse(courseDTO)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, course)
}
/* 
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
}  
func (cc *CourseController) GetCourses(c *gin.Context) {
    courses, apiErr := cc.Service.GetCourses()
    if apiErr != nil {
        c.JSON(apiErr.Status(), apiErr)
        return
    }

    c.JSON(http.StatusOK, courses)
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
} */

/* func (s *CourseService) SearchCourse(searchQuery string) ([]*domain_courses.Course, error) {
    var courses []*domain_courses.Course
    // Realiza la búsqueda en la base de datos utilizando el DB del servicio
    // Aquí deberías implementar la lógica específica de tu base de datos y la consulta para buscar cursos
    
    // Por ejemplo, podrías usar un método Find de GORM si estás usando una base de datos compatible con GORM
    if err := s.DB.Where("name LIKE ?", "%"+searchQuery+"%").Find(&courses).Error; err != nil {
        return nil, err
    }

    return courses, nil
} */



// OTRA LOGICA QUE NO IMPORTA
/* // controllers/course_controller.go
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
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course ID"})
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
} */

/* package controllers

import (
	"backend/domain/courses"
	"backend/domain/results"
	courseServices "backend/services/courses"
	"net/http"
	"strings"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	query := strings.TrimSpace(c.Query("query"))
	searchResults, err := courseServices.Search(query) // Usa el alias `courseServices` aquí
	if err != nil {
		c.JSON(http.StatusInternalServerError, results.Result{ // Usa el paquete correcto para Result
			Message: fmt.Sprintf("Error in search: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, courses.SearchResponse{ // Usa el paquete correcto para SearchResponse
		Results: searchResults,
	})
}

func Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, results.Result{ // Usa el paquete correcto para Result
			Message: fmt.Sprintf("Invalid ID: %s", err.Error()),
		})
		return
	}

	course, err := courseServices.Get(id) // Usa el alias `courseServices` aquí
	if err != nil {
		c.JSON(http.StatusNotFound, results.Result{ // Usa el paquete correcto para Result
			Message: fmt.Sprintf("Error in get: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, course)
}
package controllers

import (
	"backend/domain/courses"
	"backend/domain/results"
	courseServices "backend/services/courses"
	"net/http"
	"strings"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	query := strings.TrimSpace(c.Query("query"))
	searchResults, err := courseServices.Search(query) // Usa el alias `courseServices` aquí
	if err != nil {
		c.JSON(http.StatusInternalServerError, results.Result{ // Usa el paquete correcto para Result
			Message: fmt.Sprintf("Error in search: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, courses.SearchResponse{ // Usa el paquete correcto para SearchResponse
		Results: searchResults,
	})
}

func Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, results.Result{ // Usa el paquete correcto para Result
			Message: fmt.Sprintf("Invalid ID: %s", err.Error()),
		})
		return
	}

	course, err := courseServices.Get(id) // Usa el alias `courseServices` aquí
	if err != nil {
		c.JSON(http.StatusNotFound, results.Result{ // Usa el paquete correcto para Result
			Message: fmt.Sprintf("Error in get: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, course)
}
 */
/* func Subscribe(c *gin.Context) {
	var request domain.SubscribeRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Result{
			Message: fmt.Sprintf("Invalid request: %s", err.Error()),
		})
		return
	}

	if err := services.Subscribe(request.UserID, request.CourseID); err != nil {
		c.JSON(http.StatusConflict, domain.Result{
			Message: fmt.Sprintf("Error in subscribe; %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, domain.Result{
		Message: fmt.Sprintf("User %d successfully subscribed to course %d", request.UserID, request.CourseID),
	})
} */


