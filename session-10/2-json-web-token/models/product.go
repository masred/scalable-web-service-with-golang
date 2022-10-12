package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title       string `json:"title" form:"title" valid:"required~Title of your product is required"`
	Description string `json:"description" form:"description" valid:"required~Description of your product is required"`
	UserID      uint   `json:"user_id"`
	User        *User  `json:"user"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	return err
}

func (p *Product) BeforeUpdate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	return err
}
