package main

import (
	"sesi10/config"
	"sesi10/controllers"
	"sesi10/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()

	userController := controllers.NewUserController(db)
	productController := controllers.NewProductController(db)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userController.UserRegister)
		userRoute.POST("/login", userController.UserLogin)
	}

	productRoute := route.Group("/products")
	{
		productRoute.Use(middlewares.Auth())
		productRoute.POST("/", productController.CreateProduct)
		productRoute.GET("/", productController.GetProducts)
	}

	route.Run(config.APP_PORT)

}
