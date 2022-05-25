package routers

import (
	"sesi6-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/cars", controllers.GetCars)
	router.GET("/cars/:carId", controllers.GetCarById)
	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carId", controllers.UpdateCarById)
	router.DELETE("/cars/:carId", controllers.DeleteCarById)

	return router
}
