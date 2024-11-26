package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"unique;not null" json:"username"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"not null" json:"-"`
	AvatarURL string `json:"avatarURL"`
	Role      string `gorm:"default:'participant'" json:"role"`
}
