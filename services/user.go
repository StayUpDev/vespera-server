package services

import (
	"fmt"
	"vespera-server/models"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}

func AddUserImage(db *gorm.DB, userID string, imageURL string) error {

	if !EventExists(db, userID) {
		return fmt.Errorf("user with ID %s does not exist", userID)

	}

	return db.Model(&models.User{}).Where("id = ?", userID).Update("avatar_url", imageURL).Error
}
