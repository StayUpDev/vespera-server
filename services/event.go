package services

import (
	"fmt"
	"log"
	"vespera-server/models"

	"gorm.io/gorm"
)

func GetAllEvents(db *gorm.DB) ([]models.Evento, error) {
	var events []models.Evento
	err := db.Find(&events).Error
	return events, err
}

func GetEventsByUserID(db *gorm.DB, userID string) ([]models.Evento, error) {
	var events []models.Evento
	if !UserExists(db, userID) {
		 log.Printf("Event with ID %s does not exist", userID)
		return events, fmt.Errorf("event with ID %s does not exist", userID)
	}
	err := db.Where("user_id = ?", userID).Find(&events).Error
	
	if(err != nil){
		log.Printf("Error: %s", err)
	}
	
	return events, err
}

func GetEventByID(db *gorm.DB, id uint) (models.Evento, error) {

	var event models.Evento
	if !EventExists(db, id) {
		log.Printf("Event with ID %d does not exist", id)
		return event, fmt.Errorf("event with ID %d does not exist", id)
	}

	err := db.Preload("EventoLikes").Preload("EventoCommento").First(&event, id).Error
	return event, err
}
func CreateEvent(db *gorm.DB, event *models.Evento) error {
	return db.Create(event).Error
}
func UpdateEvent(db *gorm.DB, id string, updatedEvent models.Evento) error {
	return db.Model(&models.Evento{}).Where("id = ?", id).Updates(updatedEvent).Error
}
func DeleteEventByID(db *gorm.DB, id uint) error {
	if !EventExists(db, id) {
		return fmt.Errorf("event with ID %d does not exist", id)
	}
	return db.Delete(&models.Evento{}, id).Error
}
func DeleteEventsByUserID(db *gorm.DB, userID string) error {
	return db.Where("user_id = ?", userID).Delete(&models.Evento{}).Error
}

func AddLike(db *gorm.DB, userID uint, eventoID uint) error {

	if !EventExists(db, eventoID) {
		log.Printf("Event with ID %d does not exist", eventoID)
		return fmt.Errorf("event with ID %d does not exist", eventoID)
	}
	var existingLike models.EventoLike
	if err := db.Where("user_id = ? AND evento_id = ?", userID, eventoID).First(&existingLike).Error; err == nil {
		return nil
	}

	eventLike := models.EventoLike{
		UserID:   userID,
		EventoID: eventoID,
	}

	if err := db.Create(&eventLike).Error; err != nil {
		return err
	}

	return nil
}

func RemoveLike(db *gorm.DB, userID uint , eventoID uint) error {

	if !EventExists(db, eventoID) {
		return fmt.Errorf("event with ID %d does not exist", eventoID)
	}
	if err := db.Where("user_id = ? AND evento_id = ?", userID, eventoID).Delete(&models.EventoLike{}).Error; err != nil {
		return err
	}

	return nil
}
func AddComment(db *gorm.DB, userID uint, eventoID uint, content string) error {
	eventoCommento := models.EventoCommento{
		UserID:   userID,
		EventoID: eventoID,
		Content:  content,
	}

	if !EventExists(db, eventoID) {
		return fmt.Errorf("event with ID %d does not exist", eventoID)

	}

	if err := db.Create(&eventoCommento).Error; err != nil {
		return err
	}

	return nil
}

func AddEventoImage(db *gorm.DB, eventoID uint, imageURL string) error {
	eventoImage := models.EventoImage{
		EventoID: eventoID,
		Url:      imageURL,
	}

	if !EventExists(db, eventoID) {
		return fmt.Errorf("event with ID %d does not exist", eventoID)

	}

  fmt.Printf("Creating evento with url %s and id %d", imageURL, eventoID)

	if err := db.Create(&eventoImage).Error; err != nil {
		return err
	}

	return nil
}

// check if event exists
func EventExists(db *gorm.DB, id uint) bool {
	var count int64
	db.Model(&models.Evento{}).Where("id = ?", id).Count(&count)
	return count > 0
}
