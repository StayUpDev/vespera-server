package models

import (
	"gorm.io/gorm"
)

type EventoLike struct {
	gorm.Model
	UserID    uint `gorm:"type:varchar(36);not null" json:"userID"`
	EventoID  uint		`gorm:"type:varchar(36);not null" json:"eventoID"`

}

func (EventoLike) TableName() string {
	return "event_likes"
}
