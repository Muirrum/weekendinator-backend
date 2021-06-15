package main

import (
	"log"
	"os"

	"github.com/Muirrum/weekendinator-backend/api"
	"github.com/Muirrum/weekendinator-backend/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	log.Print("Loading env")
	err := godotenv.Load()
	log.Print("Connecting to database")
	err = db.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	if dev := os.Getenv("IS_DEV"); dev == "" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	api_group := r.Group("/api")
	api.SetupRoutes(api_group)

	// Health Check
	r.GET("/health", healthCheck)

	err = r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"data": "OK"})
}
