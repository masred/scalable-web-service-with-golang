package domain

import "time"

type SocialMedia struct {
	ID             string    `gorm:"primaryKey;type:VARCHAR(100)" json:"id"`
	Name           string    `gorm:"not null" validate:"required" json:"name"`
	SocialMediaUrl string    `gorm:"not null" validate:"required" json:"social_media_url"`
	UserID         string    `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}
