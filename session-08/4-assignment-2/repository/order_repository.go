package repository

import (
	"github.com/masred/scalable-web-service-with-golang/session-08/4-assignment-2/model/domain"
)

type OrderRepository interface {
	Create(order domain.Order) error
	Update(order domain.Order, id int) error
	Delete(id int) error
	FindById(id int) (domain.Order, error)
	FindAll() ([]domain.Order, error)
}
