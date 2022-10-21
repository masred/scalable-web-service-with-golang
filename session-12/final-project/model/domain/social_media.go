package domain

import "github.com/masred/scalable-web-service-with-golang/session-12/final-project/model"

type SocialMedia struct {
	model.Gorm
	Name           string `gorm:"not null" validate:"required" json:"name"`
	SocialMediaUrl string `gorm:"not null" validate:"required" json:"social_media_url"`
	UserID         string `json:"user_id"`
}
