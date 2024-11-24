package main

import (
	"log"
	"vespera-server/bucket"
	"vespera-server/database"
	"vespera-server/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.Println("Loaded environment variables")

	database.Connect()
	bucket.Setup()

	r := gin.Default()

	r.POST("/api/users/create", handlers.Register)
	r.POST("/api/users/validate", handlers.Login)

	r.POST("/api/events/create", handlers.CreateEventHandler)
	r.GET("/api/events/all", handlers.GetAllEventsHandler)
	r.GET("/api/events/user", handlers.GetEventsByUserIDHandler)
	r.GET("/api/events", handlers.GetEventByIDHandler)
	r.POST("/api/events/update", handlers.UpdateEventHandler)
	r.GET("/api/events/delete", handlers.DeleteEventHandler)
	r.POST("/api/evento/images/upload", handlers.UploadEventoImage)
	r.POST("/api/users/images/upload", handlers.UploadUserImage)

	r.Run(":8080")
}