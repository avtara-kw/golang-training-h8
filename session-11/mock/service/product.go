package service

import (
	"errors"
	"mock/entity"
	"mock/repository"
)

type ProductService struct {
	productRepo repository.ProductRepo
}

func (p *ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := p.productRepo.FindByID(id)
	if product == nil {
		return nil, errors.New("there is no product with id " + id)
	}

	return product, nil
}

func (p *ProductService) GetAllProducts() (*[]entity.Product, error) {
	product := p.productRepo.FindAll()
	if product == nil {
		return nil, errors.New("there is no products")
	}

	return product, nil
}
