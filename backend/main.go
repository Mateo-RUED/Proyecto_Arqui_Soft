package main

import (
	"backend/db"
	"backend/router"
    "log"

	"github.com/gin-gonic/gin"
)

func main() { //Unica funcion es inicializar
    r := gin.Default()

	db.Init()

	engine := gin.New()
	router.MapUrls(engine)
	engine.Run(":8080")
    // Configurar rutas 
    router.SetupRoutes(r, db)
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("failed to run server: %v", err)
    }

}


/* package main

import (
    "Proyecto_Arqui_Soft/backend/routers"
    "Proyecto_Arqui_Soft/backend/database"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func main() {
    r := gin.Default()

    // Conectar a la base de datos
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Inicializar la base de datos
    database.Setup(db)

    // Configurar rutas
    routers.SetupRoutes(r, db)

    // Ejecutar el servidor
    r.Run(":8080")
} */

