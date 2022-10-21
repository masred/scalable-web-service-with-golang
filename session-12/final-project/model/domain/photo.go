package domain

import "github.com/masred/scalable-web-service-with-golang/session-12/final-project/model"

type Photo struct {
	model.Gorm
	Title    string `gorm:"not null" validate:"required" json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `gorm:"not null" validate:"required" json:"photo_url"`
	UserID   string `json:"user_id"`
}
