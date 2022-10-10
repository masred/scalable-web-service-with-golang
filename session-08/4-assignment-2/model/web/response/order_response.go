package response

import (
	"time"

	"github.com/masred/scalable-web-service-with-golang/session-08/4-assignment-2/model/domain"
)

type OrderResponse struct {
	ID           uint          `json:"id"`
	CustomerName string        `json:"customer_name"`
	OrderedAt    time.Time     `json:"ordered_at"`
	Items        []domain.Item `json:"items"`
}
