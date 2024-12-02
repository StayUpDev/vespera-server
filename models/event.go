package models

import (
	"time"

	"gorm.io/gorm"
)

type Evento struct {
	gorm.Model
	Label       string    `gorm:"type:varchar(255);not null" json:"label"`
	Description string    `gorm:"type:text;not null" json:"description"`
	DateFrom    time.Time `gorm:"not null" json:"dateFrom"`
	DateTo      time.Time `gorm:"not null" json:"dateTo"`
	Category    string    `gorm:"type:varchar(100);not null" json:"category"`
	Costo       float64   `gorm:"not null" json:"costo"`
	UserID      string    `gorm:"type:varchar(36);not null" json:"userID"`
	Parcheggio  bool      `gorm:"not null" json:"parcheggio"`
	DressCode   *string   `gorm:"type:varchar(255)" json:"dressCode,omitempty"`
	Tags        []string  `gorm:"type:text[]" json:"tags"`
	Thumbnail   string    `gorm:"type:varchar(255);not null" json:"thumbnail"`

	EventLikes    []EventoLike     
	EventCommenti []EventoCommento 
}

func (Evento) TableName() string {
	return "eventi" // Specify the name of your table in the DB
}
