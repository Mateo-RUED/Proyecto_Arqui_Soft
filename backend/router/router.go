package router

import (
	"backend/controllers/users"
	"backend/controllers/courses"
	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {

	engine.POST("/users/login", users.Login)
	engine.POST("/users/register", users.CreateUser) 

}

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    // Crear instancias de servicio y controlador
    courseService := &services.CourseService{DB: db}
    courseController := &controllers.CourseController{Service: courseService}

    // Definir las rutas y asignar los métodos del controlador
    r.POST("/courses", courseController.CreateCourse)
    r.PUT("/courses/:id", courseController.UpdateCourse)
    r.DELETE("/courses/:id", courseController.DeleteCourse)
}
