package repositories

import "latihan/models"

type ProductRepo interface {
	CreateProduct(model *models.Product) error
	GetAllProducts() (*[]models.Product, error)
	GetProductByBrand(brand string) (*[]models.Product, error)
}

type productRepo struct {
	products *[]models.Product
}

func NewProductRepo(products *[]models.Product) ProductRepo {
	return &productRepo{
		products: products,
	}
}

func (p *productRepo) CreateProduct(model *models.Product) error {
	*p.products = append(*p.products, *model)

	return nil
}

func (p *productRepo) GetAllProducts() (*[]models.Product, error) {
	return p.products, nil
}

func (p *productRepo) GetProductByBrand(brand string) (*[]models.Product, error) {
	var products []models.Product
	for _, product := range *p.products {
		if product.Brand == brand {
			products = append(products, product)
		}
	}

	return &products, nil
}
