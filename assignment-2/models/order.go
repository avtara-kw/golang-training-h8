package models

import "time"

type Order struct {
	OrderID      int
	CustomerName string
	OrderedAt    *time.Time
	Items        []Item
}

type Item struct {
	ItemID      int
	ItemCode    string
	Description string
	Quantity    int
	OrderID     int
}
