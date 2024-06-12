package router

import (
    "backend/controllers/courses"
    "backend/controllers/users"
    "backend/controllers/inscripciones"
	"github.com/gin-gonic/gin"
)


func MapUrls(engine *gin.Engine) {

    users := engine.Group("/users")
    {
	users.POST("/login", controller_users.Login)
	users.POST("/register", controller_users.CreateUser) 
    }

    courses := engine.Group("/courses")
    {
    courses.POST("/create_course", controller_courses.CreateCourse)
    courses.DELETE("/delete_course", controller_courses.DeleteCourseName)
    courses.PUT("/edit_course", controller_courses.EditCourse)
    courses.GET("/get_course", controller_courses.GetCourseByID)
    courses.GET("/:id", controller_courses.GetCourseByIDParam)
    courses.DELETE("/:id", controller_courses.DeleteCourseByIDParam)
    courses.GET("/category", controller_courses.GetCoursesByCategory)

    }
    inscripciones := engine.Group("/inscripciones")
    {
        inscripciones.POST("/inscribir", controller_inscripciones.InscribirUsuario)  // Nueva ruta de inscripción
        inscripciones.GET("/users/:usuarioID/courses", controller_inscripciones.ListCoursesByUser)
    }
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










