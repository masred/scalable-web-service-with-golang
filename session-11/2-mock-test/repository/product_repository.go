package repository

import "github.com/masred/scalable-web-service-with-golang/session-11/2-mock-test/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
