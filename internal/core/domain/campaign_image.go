package domain

import "gorm.io/gorm"

// CampaignImage is an entity that represents a campaign image
type CampaignImage struct {
	gorm.Model
	CampaignID    int
	ImageFilename string
	IsPrimary     bool
}
