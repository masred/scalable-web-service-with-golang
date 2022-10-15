package service

import (
	"errors"

	"github.com/masred/scalable-web-service-with-golang/session-11/1-unit-test/entity"
	"github.com/masred/scalable-web-service-with-golang/session-11/1-unit-test/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}
