package main

import (
	"log"
	"os"
	"github.com/MohamadAlturky/Resources/core/db"
	"github.com/MohamadAlturky/Resources/core/routes"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}


	db.InitMongoDB()


	router := routes.SetupRouter()


	port := os.Getenv("PORT")
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
