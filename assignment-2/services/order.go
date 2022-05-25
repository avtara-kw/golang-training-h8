package services

import (
	"assignment-2/models"
	"assignment-2/params"
	"assignment-2/repositories"
	"database/sql"
)

type OrderService struct {
	orderRepo repositories.OrderRepo
}

var db = sql.DB{}
var repo = repositories.NewOrderRepo(&db)

func NewOrderService(repo repositories.OrderRepo) *OrderService {
	return &OrderService{
		orderRepo: repo,
	}
}

func (p *OrderService) CreateOrder(request params.Order) *params.Response {
	var err error

	order, err := p.orderRepo.CreateOrder(request.CustomerName, request.OrderedAt)
	if err != nil {
		return &params.Response{
			Status:         400,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	for _, item := range request.Items {
		err := p.orderRepo.CreateItem(&models.Item{ItemCode: item.ItemCode, Description: item.Description,
			Quantity: item.Quantity, OrderID: order.OrderID})
		if err != nil {
			return &params.Response{
				Status:         400,
				Error:          "BAD REQUEST",
				AdditionalInfo: err.Error(),
			}
		}

	}

	return &params.Response{
		Status:  200,
		Message: "CREATE SUCCESS",
		Payload: order,
	}
}

func (p *OrderService) GetAllOrder() *params.Response {
	var err error

	order, err := p.orderRepo.GetAllOrder()
	if err != nil {
		return &params.Response{
			Status:         400,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  200,
		Message: "GET DATA SUCCESS",
		Payload: order,
	}
}
