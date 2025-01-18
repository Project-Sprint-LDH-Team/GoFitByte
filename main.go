package main

import (
	"fit-byte-go/pkg/database"
	"log"
)

func main() {
	_, err := database.Init()
	if err != nil {
		log.Fatalf("gagal: %v\n", err)
	}

	// r := gin.Default()

	// r.GET("/api/v1/products", handlers.ListProducts(db))
}
