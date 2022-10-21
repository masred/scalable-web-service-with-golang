package domain

import "github.com/masred/scalable-web-service-with-golang/session-12/final-project/model"

type Comment struct {
	model.Gorm
	UserID  string `gorm:"type:VARCHAR(50);not null" json:"user_id"`
	PhotoID string `gorm:"type:VARCHAR(50);not null" json:"photo_id"`
	Message string `gorm:"not null" validate:"required" json:"message"`
}
