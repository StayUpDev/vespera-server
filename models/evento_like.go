package models

import (
	"time"
)

type EventoLike struct {
	ID        string    `gorm:"primaryKey;type:varchar(36);not null" json:"$id"`
	UserID    string    `gorm:"type:varchar(36);not null" json:"userID"`
	EventoID  string    `gorm:"type:varchar(36);not null" json:"eventoID"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (EventoLike) TableName() string {
	return "event_likes"
}
