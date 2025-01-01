package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User is an entity that represents a user
type User struct {
	gorm.Model
	ID             uuid.UUID `gorm:"primaryKey;type:char(36);not null"`
	Name           string    `gorm:"type:varchar(100);not null"`
	Occupation     string    `gorm:"type:varchar(100)"`
	Email          string    `gorm:"type:varchar(50);not null;unique"`
	PasswordHash   string    `gorm:"type:varchar(255);not null"`
	AvatarFilename string    `gorm:"type:varchar(255)"`
	Role           string    `gorm:"type:enum('user','admin');default:'user'"`
	Token          string    `gorm:"type:varchar(255)"`
}
