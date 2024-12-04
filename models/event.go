package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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
	Parcheggio  bool      `json:"parcheggio"`
	DressCode   *string   `gorm:"type:varchar(255)" json:"dressCode,omitempty"`
	Thumbnail   string    `gorm:"type:varchar(255);" json:"thumbnail"`

	EventLikes    []EventoLike     
	EventCommenti []EventoCommento 
}

func (Evento) TableName() string {
	return "eventi" // Specify the name of your table in the DB
}
type StringArray []string

// Implement the Scanner interface for reading from the DB
func (s *StringArray) Scan(value interface{}) error {
    strValue, ok := value.(string)
    if !ok {
        return fmt.Errorf("unsupported type for StringArray: %T", value)
    }

    return json.Unmarshal([]byte(strValue), s)
}

// Implement the Valuer interface for writing to the DB
func (s StringArray) Value() (driver.Value, error) {
    return json.Marshal(s)
}
