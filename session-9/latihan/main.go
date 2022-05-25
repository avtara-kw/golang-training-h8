package main

import (
	"latihan/models"
	"latihan/repositories"
	"latihan/server"
	"latihan/server/controllers"
	"latihan/server/middleware"
	"latihan/services"
)

/*
	Latihan :
	Buatlah sebuah REST API untuk :
		- Register
		  ini cukup dengan menginputkan username dan password.
		  data user bisa lebih dari 1, jadi silahkan gunakan slice.

		- GetAllUsers
		  untuk cek seluruh user. endpoint ini khusus untuk user yg sudah
		  didaftarkan.

		- AddNewProducts
		  yang bisa hit endpoint ini hanyalah user yang sudah di daftarkan.
		  data yang dibutuhkan adalah :
		  	- Nama
			- Brand
			- Stok
			- Price

		- GetProducts
		  tidak ada auth disini. jadi ini adalah API open. akan ngereturn :
		  payload : [
			  {
				  nama 	: "",
				  brand : "",
				  stok 	: 0,
				  price	: 0
			  }
		  ]

		- GetProductByBrand
		  tidak ada auth disini. jadi ini adalah API open.
		  Brand akan di dapat dari query. akan ngereturn :
		  payload : [
			  {
				  nama 	: "",
				  brand : "",
				  stok 	: 0,
				  price	: 0
			  }
		  ]
*/

func main() {
	middleware := middleware.NewMiddleware(&models.Users)

	userRepo := repositories.NewUserRepo(&models.Users)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService, middleware)

	productRepo := repositories.NewProductRepo(&models.Products)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService, middleware)

	server := server.NewServer(userController, productController)
	server.StartServer()
}
