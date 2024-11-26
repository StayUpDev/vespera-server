package models

import (
	"gorm.io/gorm"
)

type EventoCommento struct {
	gorm.Model
	EventoID  uint 		`json:"eventoID"`
	UserID  uint 		`json:"userID"`
	Content   string    `gorm:"type:text;not null" json:"content"`
}

// TableName specifies the name of the table
func (EventoCommento) TableName() string {
	return "evento_commenti" // Name of the table for comments
}
