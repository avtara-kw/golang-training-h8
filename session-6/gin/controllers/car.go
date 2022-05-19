package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarId string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

type PayloadSuccess struct {
	Status  int         `json:"status"`
	Message string      `json:"msg,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

var Cars = []Car{}

func CreateCar(ctx *gin.Context) {
	var newCar Car

	err := ctx.ShouldBindJSON(&newCar)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newCar.CarId = fmt.Sprintf("car-%d", len(Cars)+1)

	Cars = append(Cars, newCar)
	ctx.JSON(http.StatusCreated, PayloadSuccess{
		Status:  http.StatusCreated,
		Message: "CREATED SUCCESS",
	})
}

func GetCars(ctx *gin.Context) {
	if len(Cars) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": http.StatusNotFound,
			"msg":   "NO DATA",
		})
		return
	}
	ctx.JSON(http.StatusOK, PayloadSuccess{
		Status:  http.StatusOK,
		Payload: Cars,
	})
}

func GetCarById(ctx *gin.Context) {
	carId := ctx.Param("carId")
	condition := false

	var car Car

	for _, c := range Cars {
		if c.CarId == carId {
			car = c
			condition = true
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": http.StatusNotFound,
			"msg":   "NOT FOUND",
		})
		return
	}

	ctx.JSON(http.StatusOK, PayloadSuccess{
		Status:  http.StatusOK,
		Payload: car,
	})
}

func UpdateCarById(ctx *gin.Context) {
	carId := ctx.Param("carId")
	isExist := false

	var updateCar Car

	if err := ctx.ShouldBindJSON(&updateCar); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":           http.StatusBadRequest,
			"msg":             "BAD REQUEST",
			"additional_info": err.Error(),
		})
		return
	}

	for i, c := range Cars {
		if c.CarId == carId {
			updateCar.CarId = c.CarId
			Cars[i] = updateCar
			isExist = true
			break
		}
	}

	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": http.StatusNotFound,
			"msg":   "NOT FOUND",
		})
		return
	}

	ctx.JSON(http.StatusOK, PayloadSuccess{
		Status:  http.StatusOK,
		Message: "UPDATED SUCCESS",
	})
}

func DeleteCarById(ctx *gin.Context) {
	carId := ctx.Param("carId")
	isExist := false
	index := 0

	for i, c := range Cars {
		if c.CarId == carId {
			isExist = true
			index = i
			break
		}
	}

	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": http.StatusNotFound,
			"msg":   "NOT FOUND",
		})
		return
	}

	cars := append(Cars[:index], Cars[index+1:]...)
	Cars = cars
	ctx.JSON(http.StatusOK, PayloadSuccess{
		Status:  http.StatusOK,
		Message: "DELETED SUCCESS",
		Payload: Cars,
	})
}
