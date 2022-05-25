package controllers

import (
	"assignment-2/params"
	"assignment-2/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{
		orderService: *service,
	}
}

func (p *OrderController) CreateNewOrder(c *gin.Context) {
	var req params.Order

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		})
		return
	}

	response := p.orderService.CreateOrder(req)
	c.JSON(response.Status, response)
}

func (p *OrderController) GetAllOrder(c *gin.Context) {
	response := p.orderService.GetAllOrder()
	c.JSON(response.Status, response)
}
