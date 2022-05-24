package models

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	OrderID      int `gorm:"primary_key"`
	CustomerName string
	Items        []Items `gorm:"ForeignKey:OrderID"`
}

type Items struct {
	gorm.Model
	ItemID      int `gorm:"primary_key"`
	ItemCode    int
	Description string
	Quantity    int
	Order_id    int
}
