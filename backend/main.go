package main

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

}


