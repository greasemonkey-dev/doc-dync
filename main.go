package main

import (
	"doc-sync/handlers"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting working directory: ", err)
	}
	log.Println("Current working directory:", wd)
	// Retrieve the port from the environment variables
	err = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not found in .env file")
	}
	r := gin.Default()

	r.GET("/appointments", handlers.ParseHandler)

	log.Println("Server is running...")
	err = r.Run(port)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
