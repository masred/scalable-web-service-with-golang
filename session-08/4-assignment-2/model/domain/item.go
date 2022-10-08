package domain

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint
}
