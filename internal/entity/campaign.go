package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Campaign is an entity that represents a campaign
type Campaign struct {
	gorm.Model
	ID               uuid.UUID `gorm:"primaryKey;type:char(36)"`
	UserID           uuid.UUID `gorm:"type:char(36)"`
	Name             string    `gorm:"type:varchar(255)"`
	ShortDescription string    `gorm:"type:varchar(255)"`
	Description      string    `gorm:"type:text"`
	GoalAmount       int       `gorm:"type:bigint"`
	CurrentAmount    int       `gorm:"type:bigint"`
	BackerCount      uint32    `gorm:"type:int"`
	Perks            string    `gorm:"type:text"`
	Slug             string    `gorm:"type:varchar(255)"`
}

type Campaigns []Campaign
