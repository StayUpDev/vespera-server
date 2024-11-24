package database

import (
	"fmt"
	"log"
	"os"

	"vespera-server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, name, port, sslmode)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Database connected successfully!")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	err = DB.AutoMigrate(&models.Evento{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	err = DB.AutoMigrate(&models.EventoCommento{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	err = DB.AutoMigrate(&models.EventoImage{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	err = DB.AutoMigrate(&models.EventoImage{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
