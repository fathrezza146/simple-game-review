package models

import "time"

type (
	Reviews struct {
		ID           uint           `json:"id" gorm:"primary_key"`
		GamesID      uint           `json:"game_id"`
		TextReview   string         `json:"review" gorm:"type:text"`
		CreatedAt    time.Time      `json:"created_at"`
		UpdatedAt    time.Time      `json:"updated_at"`
		Games        Games          `json:"-"`
		ReviewRating []ReviewRating `json:""`
		User         []User         `json:"-"`
	}
)
