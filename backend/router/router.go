package router

import (
	controller_archivos "backend/controllers/archivos" // Controlador de archivos
	controller_comments "backend/controllers/comments"
	controller_courses "backend/controllers/courses"
	controller_inscripciones "backend/controllers/inscripciones"
	controller_users "backend/controllers/users"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {
	// Rutas de usuarios
	users := engine.Group("/users")
	{
		users.POST("/login", controller_users.Login)
		users.POST("/register", controller_users.CreateUser)
	}

	// Rutas de cursos
	courses := engine.Group("/courses")
	{
		courses.POST("/create_course", controller_courses.CreateCourse)
		courses.DELETE("/delete_course", controller_courses.DeleteCourseName)
		courses.PUT("/edit_course", controller_courses.EditCourse)
		courses.GET("/get_course", controller_courses.GetCourseByID)
		courses.GET("/:id", controller_courses.GetCourseByIDParam)
		courses.DELETE("/:id", controller_courses.DeleteCourseByIDParam)
		courses.GET("/all", controller_courses.GetAllCourses)
		courses.GET("/category", controller_courses.GetCoursesByCategory)
	}

	// Rutas de inscripciones
	inscripciones := engine.Group("/inscripciones")
	{
		inscripciones.POST("/inscribir", controller_inscripciones.InscribirUsuario)
		inscripciones.GET("/users/:usuarioID/courses", controller_inscripciones.ListCoursesByUser)
	}

	// Rutas de comentarios
	comments := engine.Group("/comments")
	{
		comments.POST("/add_comment", controller_comments.AddComment)
		comments.GET("/:id/comments", controller_comments.GetCommentsByCourse)
	}

	// Rutas de manejo de archivos
	archivos := engine.Group("/archivos")
	{
		archivos.POST("/subir/:cursoID", controller_archivos.UploadFile)                // Subir archivos
		archivos.GET("/:cursoID/listar", controller_archivos.GetFilesByCourse)          // Listar archivos
		archivos.GET("/:cursoID/descargar/:filename", controller_archivos.DownloadFile) // Descargar archivo
	}
	
}
