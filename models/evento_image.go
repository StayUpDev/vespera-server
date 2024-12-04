package models

import (
	"gorm.io/gorm"
)

type EventoImage struct {
	gorm.Model
	EventoID  uint 		`gorm:"not null" json:"eventoID"`
	Url       string    `gorm:"type:varchar(255);not null" json:"url"`
}

func (EventoImage) TableName() string {
	return "event_image"
}
