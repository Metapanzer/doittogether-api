package domain

import "gorm.io/gorm"

// User is an entity that represents a user
type User struct {
	gorm.Model
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFilename string
	Role           string
	Token          string
}
