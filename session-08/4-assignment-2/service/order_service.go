package service

import (
	"github.com/masred/scalable-web-service-with-golang/session-08/4-assignment-2/model/domain"
	"github.com/masred/scalable-web-service-with-golang/session-08/4-assignment-2/model/web/request"
)

type OrderService interface {
	Create(request request.OrderCreate) error
	Update(request request.OrderUpdate, id int) error
	Delete(id int) error
	FindById(id int) (domain.Order, error)
	FindAll() ([]domain.Order, error)
}
