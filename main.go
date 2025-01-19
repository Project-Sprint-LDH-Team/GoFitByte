package main

import (
	"fit-byte-go/internal/handlers"
	"fit-byte-go/internal/repositories"
	"fit-byte-go/internal/services"
	"fit-byte-go/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Init()

	if err != nil {
		log.Fatalf("gagal: %v\n", err)
	}

	// init repo repostiry, service, and handler
	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	r := gin.Default()

	r.POST("/api/v1/register", authHandler.Register)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not start server %v", err)
	}
}
