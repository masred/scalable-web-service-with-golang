package domain

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}
