package domain

import "gorm.io/gorm"

// Campaign is an entity that represents a campaign
type Campaign struct {
	gorm.Model
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	GoalAmount       int
	CurrentAmount    int
	BackerCount      int
	Perks            string
	Slug             string
}
