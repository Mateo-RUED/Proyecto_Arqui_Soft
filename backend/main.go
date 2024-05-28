package main

import (
	"backend/db"
	"backend/router"

	"github.com/gin-gonic/gin"
)

func main() { //Unica funcion es inicializar

	db.Init()

	engine := gin.New()
	router.MapUrls(engine)
	engine.Run(":8080")

}

