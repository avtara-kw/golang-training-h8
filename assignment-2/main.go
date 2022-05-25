package main

import (
	"assignment-2/app/config"
	"assignment-2/controllers"
	"assignment-2/repositories"
	"assignment-2/services"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	AppPort = ":8888"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	orderRepo := repositories.NewOrderRepo(db)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)
	router.POST("/orders", orderController.CreateNewOrder)
	router.GET("/orders", orderController.GetAllOrder)

	log.Println("server running at port ", AppPort)
	router.Run(AppPort)
}
