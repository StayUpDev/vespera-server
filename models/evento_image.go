package models

import (
	"time"
)

type EventoImage struct {
	ID        string    `gorm:"primaryKey;type:varchar(36);not null" json:"$id"`
	EventoID  string    `gorm:"type:varchar(36);not null" json:"eventoID"`
	Url       string    `gorm:"type:varchar(255);not null" json:"url"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (EventoImage) TableName() string {
	return "event_image"
}
