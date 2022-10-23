package domain

import "time"

type Photo struct {
	ID        string    `gorm:"primaryKey;type:VARCHAR(100)" json:"id"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `gorm:"not null" validate:"required" json:"photo_url"`
	UserID    string    `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
