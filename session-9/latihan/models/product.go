package models

type Product struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Stok  int    `json:"stok"`
	Price int    `json:"price"`
}

var Products []Product
