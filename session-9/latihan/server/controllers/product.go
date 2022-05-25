package controllers

import (
	"encoding/json"
	"latihan/params"
	"latihan/server/middleware"
	"latihan/services"
	"net/http"
)

type ProductController struct {
	productService *services.ProductServices
	middleware     *middleware.Middleware
}

func NewProductController(productService *services.ProductServices, middleware *middleware.Middleware) *ProductController {
	return &ProductController{
		productService: productService,
		middleware:     middleware,
	}
}

func (p *ProductController) CreateProduct(rw http.ResponseWriter, r *http.Request) {
	if !p.middleware.CheckAuth(rw, r) {
		return
	}
	var req params.ProductCreate

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response := params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}

		params.WriteJsonResponse(rw, &response)
		return
	}

	response := p.productService.CreateProduct(&req)

	params.WriteJsonResponse(rw, response)
}

func (p *ProductController) GetAllProducts(rw http.ResponseWriter, r *http.Request) {
	if !p.middleware.CheckAuth(rw, r) {
		return
	}

	query := r.URL.Query()
	brand := query.Get("brand")
	var response *params.Response
	if brand == "" {
		response = p.productService.GetAllProducts()
	} else {
		response = p.productService.GetProductByBrand(brand)
	}
	params.WriteJsonResponse(rw, response)
}
