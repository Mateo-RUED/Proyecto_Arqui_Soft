/* package main

import (
	"backend/db"
	"backend/router"
    "log"

	"github.com/gin-gonic/gin"
)

func main() { //Unica funcion es inicializar
   
	db.Init()

	engine := gin.New()
	router.MapUrls(engine)
	engine.Run(":8080")
    // Configurar rutas 
    r := gin.Default()
    router.SetupRoutes(r, db)
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("failed to run server: %v", err)
    }

} */


package main

import (
	"backend/db"
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	router := gin.New()
	// Middleware to handle CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
	})
	router.GET("/courses/search", controllers.Search)
	router.GET("/courses/:id", controllers.Get)
	router.POST("/subscriptions", controllers.Subscribe)
	router.Run(":8080")
}
	



