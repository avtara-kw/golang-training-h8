package params

type ProductCreate struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Stok  int    `json:"stok"`
	Price int    `json:"price"`
}
