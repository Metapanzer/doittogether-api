package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CampaignImage is an entity that represents a campaign image
type CampaignImage struct {
	gorm.Model
	CampaignID    uuid.UUID `gorm:"type:char(36)"`
	ImageFilename string    `gorm:"type:varchar(255)"`
	IsPrimary     bool      `gorm:"type:tinyint;default:0"`
}
