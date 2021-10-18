package models

import "time"

type (
	Developers struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Games     []Games   `json:"-"`
	}
)
