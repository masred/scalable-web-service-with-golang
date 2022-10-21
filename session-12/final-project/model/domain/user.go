package domain

import "github.com/masred/scalable-web-service-with-golang/session-12/final-project/models"

type User struct {
	models.Gorm
	Username    string        `gorm:"not null;unique" validate:"required,unique"  json:"username"`
	Email       string        `gorm:"not null;unique" validate:"email,required,unique" json:"email"`
	Password    string        `gorm:"not null" validate:"required,min:8" json:"password"`
	Age         int           `json:"age"`
	SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"social_media"`
	Comment     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"comment"`
	Photo       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"photo"`
}
