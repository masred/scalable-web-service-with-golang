package service

import (
	"testing"

	"github.com/masred/scalable-web-service-with-golang/session-11/1-unit-test/entity"
	"github.com/masred/scalable-web-service-with-golang/session-11/1-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetOneProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := entity.Product{
		Id:   "2",
		Name: "Sun Glasses",
	}

	productRepository.Mock.On("FindById", "2").Return(product)

	result, err := productService.GetOneProduct("2")

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, product.Id, result.Id, "result has to be '2'")
	assert.Equal(t, product.Name, result.Name, "result has to be 'sun glasses'")
	assert.Equal(t, &product, result, "result has to be a product data with id '2'")
}
