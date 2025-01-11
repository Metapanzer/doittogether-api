package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Transaction is an entity that represents a transaction
type Transaction struct {
	gorm.Model
	ID         uuid.UUID `gorm:"primaryKey;type:char(36)"`
	UserID     uuid.UUID `gorm:"type:char(36)"`
	CampaignID uuid.UUID `gorm:"type:char(36)"`
	Amount     uint32    `gorm:"type:int"`
	Status     string    `gorm:"type:varchar(50)"`
	Code       string    `gorm:"type:varchar(100)"`
}
