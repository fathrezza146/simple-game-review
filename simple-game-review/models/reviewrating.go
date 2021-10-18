package models

import "time"

type (
	ReviewRating struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		ReviewsID uint      `json:"review_id"`
		Helpful   bool      `json:"helpful"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Reviews   Reviews   `json:"-"`
	}
)
