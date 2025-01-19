package main

import (
	"fit-byte-go/internal/handlers"
	"fit-byte-go/internal/middleware"
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
	// auth
	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// user
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()

	//main root
	g := r.Group("/v1")

	g.POST("/register", authHandler.Register)
	g.POST("/login", authHandler.Login)

	g.Use(middleware.AuthMiddleware())
	{
		g.PATCH("/user", userHandler.UpdateUser)
		g.GET("/user", userHandler.GetUser)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not start server %v", err)
	}
}
