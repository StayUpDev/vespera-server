package services

import (
	"fmt"
	"log"
	"vespera-server/models"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}

func AddUserImage(db *gorm.DB, userID string, imageURL string) error {

	if !UserExists(db, userID) {
		return fmt.Errorf("user with ID %s does not exist", userID)

	}

	return db.Model(&models.User{}).Where("id = ?", userID).Update("avatar_url", imageURL).Error
}

func GetUserByID(db *gorm.DB, id string) (models.User, error) {

	var user models.User
	if !UserExists(db, id) {
		log.Printf("user with ID %s does not exist", id)
		return user , fmt.Errorf("event with ID %s does not exist", id)
	}

	err := db.Model(&models.User{}).Where("id = ?", id).First(&user).Error
	return user , err
}

func UserExists(db *gorm.DB, id string) bool {
	var count int64
	db.Model(&models.User{}).Where("id = ?", id).Count(&count)
	return count > 0
}