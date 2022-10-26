package domain

import "time"

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	PhotoID   uint   `gorm:"not null"`
	Message   string `gorm:"not null"`
	UpdatedAt time.Time
	CreatedAt time.Time
	User      User  `gorm:"foreignKey:UserID"`
	Photo     Photo `gorm:"foreignKey:PhotoID"`
}
