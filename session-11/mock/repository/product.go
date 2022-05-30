package repository

import "mock/entity"

type ProductRepo interface {
	FindByID(id string) *entity.Product
	FindAll() *[]entity.Product
}
