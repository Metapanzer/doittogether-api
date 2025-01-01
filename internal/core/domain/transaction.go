package domain

import "gorm.io/gorm"

// Transaction is an entity that represents a transaction
type Transaction struct {
	gorm.Model
	UserID     int
	CampaignID int
	Amount     int
	Status     string
	Code       string
}
