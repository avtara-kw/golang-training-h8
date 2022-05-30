package service

import (
	"mock/entity"
	"mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repo = repository.ProductRepoMock{Mock: mock.Mock{}}
var svc = ProductService{productRepo: &repo}

func TestProductGetOneNotFound(t *testing.T) {
	repo.Mock.On("FindByID", "2").Return(nil)

	product, err := svc.GetOneProduct("2")
	assert.NotNil(t, err)
	assert.Nil(t, product)
}

func TestProductGetOne(t *testing.T) {
	id := "10"
	product := entity.Product{
		Id:   id,
		Name: "Baju",
	}

	repo.Mock.On("FindByID", id).Return(product)

	res, err := svc.GetOneProduct(id)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, product.Id, res.Id)
	assert.Equal(t, product.Name, res.Name)
}

func TestProductGetAll(t *testing.T) {

	product := []entity.Product{
		{Id: "1", Name: "Baju"},
		{Id: "2", Name: "Baju"},
	}

	repo.Mock.On("FindAll").Return(product)

	res, err := svc.GetAllProducts()
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
