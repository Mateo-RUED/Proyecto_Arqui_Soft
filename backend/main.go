package main

import (
	"backend/db"
	"backend/router"
	"log"

	domain_users "backend/domain/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	// Inicializar la base de datos
	db.Init()

	// Crear el usuario administrador si no existe
	var user domain_users.User
	if err := db.DB.Where("username = ?", "sergio").First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte("sergio123"), bcrypt.DefaultCost)
			if err != nil {
				log.Fatalf("Error hashing password: %v", err)
			}

			user = domain_users.User{
				Username: "sergio",
				Password: string(hashedPassword),
				Tipo:     "Profesor",
				Email:    "",
			}

			if err := db.DB.Create(&user).Error; err != nil {
				log.Fatalf("Error creating user: %v", err)
			}
		} else {
			log.Fatalf("Error querying user: %v", err)
		}
	}

	// Crear la instancia de Gin
	engine := gin.New()

	// Configurar CORS
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Permitir el origen de tu frontend
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "X-Auth-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Mapear URLs
	router.MapUrls(engine)

	// Ejecutar el servidor en el puerto 8080
	engine.Run(":8080")
}