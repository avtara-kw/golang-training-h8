package services

import (
	"latihan/models"
	"latihan/params"
	"latihan/repositories"
	"net/http"
)

type ProductServices struct {
	ProductRepo repositories.ProductRepo
}

func NewProductService(ProductRepo repositories.ProductRepo) *ProductServices {
	return &ProductServices{
		ProductRepo: ProductRepo,
	}
}

func (u *ProductServices) CreateProduct(req *params.ProductCreate) *params.Response {
	product := makeProductModelFromRequest(req)
	err := u.ProductRepo.CreateProduct(product)
	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Message: "CREATED SUCCESS",
	}
}

func (u *ProductServices) GetAllProducts() *params.Response {
	products, _ := u.ProductRepo.GetAllProducts()
	return &params.Response{
		Status:  http.StatusOK,
		Payload: products,
	}

}

func (u *ProductServices) GetProductByBrand(brand string) *params.Response {
	products, _ := u.ProductRepo.GetProductByBrand(brand)
	return &params.Response{
		Status:  http.StatusOK,
		Payload: products,
	}

}

func makeProductModelFromRequest(req *params.ProductCreate) *models.Product {
	return &models.Product{
		Name:  req.Name,
		Brand: req.Brand,
		Stok:  req.Stok,
		Price: req.Price,
	}
}
