package main

import (
	"backend/db"
	"backend/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la base de datos
	db.Init()

	// Crear la instancia de Gin
	engine := gin.New()

	// Configurar CORS
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Permitir el origen de tu frontend
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "X-Auth-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
	}))

	// Servir archivos estáticos (necesario para permitir descargas directas)
	engine.Static("/backend/uploads", "./backend/uploads")

	// Mapear URLs
	router.MapUrls(engine)

	// Ejecutar el servidor en el puerto 8080
	engine.Run(":8080")
}
