package domain

import (
	"time"

	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/helper"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null;uniqueIndex"`
	Email     string `gorm:"not null;uniqueIndex"`
	Password  string `gorm:"not null"`
	Age       int    `gorm:"not null"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.Password = helper.Hash(user.Password)

	return nil
}
