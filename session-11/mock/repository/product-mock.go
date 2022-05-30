package repository

import (
	"mock/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepoMock struct {
	Mock mock.Mock
}

func (pm *ProductRepoMock) FindByID(id string) *entity.Product {
	arguments := pm.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(entity.Product)
	return &product
}

func (pm *ProductRepoMock) FindAll() *[]entity.Product {
	arguments := pm.Mock.Called()

	if arguments.Get(0) == nil {
		return nil
	}

	products := arguments.Get(0).([]entity.Product)
	return &products
}
