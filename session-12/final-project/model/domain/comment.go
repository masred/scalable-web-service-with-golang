package domain

import "time"

type Comment struct {
	ID        string    `gorm:"primaryKey;type:VARCHAR(100)" json:"id"`
	UserID    string    `gorm:"type:VARCHAR(50);not null" json:"user_id"`
	PhotoID   string    `gorm:"type:VARCHAR(50);not null" json:"photo_id"`
	Message   string    `gorm:"not null" validate:"required" json:"message"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
