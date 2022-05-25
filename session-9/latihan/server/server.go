package server

import (
	"latihan/server/controllers"
	"log"
	"net/http"
)

type Server struct {
	userController    *controllers.UserController
	productController *controllers.ProductController
}

func NewServer(user *controllers.UserController, productController *controllers.ProductController) *Server {
	return &Server{
		userController:    user,
		productController: productController,
	}
}

func (s *Server) StartServer() {
	port := ":4444"

	server := new(http.Server)
	server.Addr = port

	http.HandleFunc("/users/register", s.userController.Register)
	http.HandleFunc("/users", s.userController.GetAllUsers)

	http.HandleFunc("/products/create", s.productController.CreateProduct)
	http.HandleFunc("/products", s.productController.GetAllProducts)

	log.Println("server running at port", port)

	server.ListenAndServe()

}
