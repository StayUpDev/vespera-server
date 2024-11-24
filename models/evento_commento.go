package models

import "time"

type EventoCommento struct {
	ID        string    `gorm:"primaryKey;type:varchar(36);not null" json:"$id"`
	UserID    string    `gorm:"type:varchar(36);not null" json:"userID"`
	EventoID  string    `gorm:"type:varchar(36);not null" json:"eventoID"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

// TableName specifies the name of the table
func (EventoCommento) TableName() string {
	return "evento_commenti" // Name of the table for comments
}
