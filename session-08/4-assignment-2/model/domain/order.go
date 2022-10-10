package domain

import (
	"time"
)

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `json:"items"`
}
