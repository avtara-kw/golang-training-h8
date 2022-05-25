package repositories

import (
	"assignment-2/models"
	"database/sql"
	"fmt"
)

type OrderRepo interface {
	CreateItem(item *models.Item) error
	CreateOrder(name, order_at string) (*models.Order, error)
	GetAllOrder() ([]models.Order, error)
}

type orderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) OrderRepo {
	return &orderRepo{db}
}

func (p *orderRepo) CreateItem(data *models.Item) error {
	var item = models.Item{}

	sqlStatement := `
INSERT INTO items (item_code, description, quantity, order_id)
VALUES ($1, $2, $3, $4) Returning *
`

	err := p.db.QueryRow(sqlStatement, data.ItemCode, data.Description, data.Quantity, data.OrderID).
		Scan(&item.ItemID, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderID)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(item)
	return nil
}

func (p *orderRepo) CreateOrder(name, order_at string) (*models.Order, error) {
	var (
		order = models.Order{}
		err   error
	)

	sqlStatementNoOrderDate := `
INSERT INTO orders (customer_name)
VALUES ($1) Returning *;
`
	sqlstatementWithOrderDate := `
INSERT INTO orders (customer_name, ordered_at)
VALUES ($1, $2) Returning *;
`
	if order_at != "" {
		err = p.db.QueryRow(sqlstatementWithOrderDate, name, order_at).
			Scan(&order.OrderID, &order.CustomerName, &order.OrderedAt)
	} else {
		err = p.db.QueryRow(sqlStatementNoOrderDate, name).
			Scan(&order.OrderID, &order.CustomerName, &order.OrderedAt)
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(order)
	return &order, nil
}

func (p *orderRepo) GetAllOrder() ([]models.Order, error) {
	var result []models.Order

	//sqlQuery := `SELECT i.*, array_agg(t.*)
	//FROM orders AS i JOIN items AS t ON t.order_id = i.order_id GROUP BY i.order_id`

	sqlQuery := `SELECT * FROM orders`

	rows, err := p.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var order models.Order
		err = rows.Scan(&order.OrderID, &order.CustomerName, &order.OrderedAt)
		if err != nil {
			return nil, err
		}
		sqlItems := `SELECT * FROM items WHERE order_id = $1`
		rowsItems, err := p.db.Query(sqlItems, order.OrderID)
		if err != nil {
			return nil, err
		}

		defer rowsItems.Close()
		for rowsItems.Next() {
			var item models.Item
			err = rowsItems.Scan(&item.ItemID, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderID)
			if err != nil {
				return nil, err
			}
			order.Items = append(order.Items, item)
		}
		result = append(result, order)
	}

	return result, nil
}
