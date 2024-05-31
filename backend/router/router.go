package router

import (
    "backend/controllers/courses"
    "backend/controllers/users"
	"backend/services/courses"
    "backend/services/users"
	"backend/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func MapUrls(engine *gin.Engine) {

	engine.POST("/users/login", users.Login)
	engine.POST("/users/register", users.CreateUser) 

}
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    // Crear instancias de servicio y controlador
    courseService := &services.CourseService{DB: db}
    courseController := &controllers.CourseController{Service: courseService} 
    r.POST("/courses", courseController.CreateCourse)
    r.PUT("/courses/:id", courseController.UpdateCourse)
    r.DELETE("/courses/:id", courseController.DeleteCourse)
} 
 
/*  // router/router.go

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    courseService := services.NewCourseService(db)
    courseController := controllers.NewCourseController(courseService)

    r.POST("/courses", courseController.CreateCourse)
    r.PUT("/courses/:id", courseController.UpdateCourse)
    r.DELETE("/courses/:id", courseController.DeleteCourse)
    r.GET("/courses/:id", courseController.GetCourseById)
    r.GET("/courses", courseController.GetCourses)
}  */

/* router.GET("/courses/search", controllers.Search)
	router.GET("/courses/:id", controllers.Get)
	router.POST("/subscriptions", controllers.Subscribe)
	router.Run(":8080") */










