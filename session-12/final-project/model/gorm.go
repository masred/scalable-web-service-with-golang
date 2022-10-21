package model

import "time"

type Gorm struct {
	ID        string    `gorm:"primaryKey;type:VARCHAR(100)" json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
