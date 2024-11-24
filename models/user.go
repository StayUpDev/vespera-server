package models

type User struct {
	ID        string `gorm:"primaryKey;type:varchar(36);not null" json:"$id"`
	Username  string `gorm:"unique;not null" json:"username"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"not null" json:"-"`
	AvatarURL string `json:"avatarURL"`
	Role      string `gorm:"default:'participant'" json:"role"`
}
