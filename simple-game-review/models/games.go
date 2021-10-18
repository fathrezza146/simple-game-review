package models

import "time"

type (
	Games struct {
		ID           uint       `gorm:"primary_key" json:"id"`
		Name         string     `json:"name"`
		Year         int        `json:"year"`
		PublishersID uint       `json:"publisher_id"`
		DevelopersID uint       `json:"dev_id"`
		CategoryID   []uint     `json:"category_id" gorm:"type:text[]"`
		CreatedAt    time.Time  `json:"created_at"`
		UpdatedAt    time.Time  `json:"updated_at"`
		Publishers   Publishers `json:"-"`
		Developers   Developers `json:"-"`
		Category     Category   `json:"-"`
		Reviews      []Reviews  `json:"-"`
	}
)
