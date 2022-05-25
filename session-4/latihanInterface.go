package main

import (
	"fmt"
	"strings"
)

/*
	buatlah struct :
	User 	=> name, email, address
	Product => name, stock, price

	buatlah interface :
	save() bool
	print()
*/

var users []User
var products []Product

type Repository interface {
	Save() bool
	Print()
}

type User struct {
	Name    string
	Email   string
	Address string
}

func NewUser(name, email, address string) Repository {
	return &User{
		Name:    name,
		Email:   email,
		Address: address,
	}
}

func (u *User) Save() bool {
	if u == nil {
		return false
	}

	if u.Name == "" || u.Address == "" || u.Email == "" {
		return false
	}

	users = append(users, *u)
	return true
}

func (u *User) Print() {
	fmt.Println(strings.Repeat("=", 20))
	for _, user := range users {
		fmt.Println("Name :", user.Name)
		fmt.Println("Email :", user.Email)
		fmt.Println("Address :", user.Address)
		fmt.Println()
	}
}

type Product struct {
	Name  string
	Stock int8
	Price int
}

func NewProduct(name string, stock int8, price int) Repository {
	return &Product{
		Name:  name,
		Stock: stock,
		Price: price,
	}
}

func (p *Product) Save() bool {
	if p == nil {
		return false
	}

	if p.Name == "" || p.Price == 0 || p.Stock <= 0 {
		return false
	}

	products = append(products, *p)
	return true
}

func (p *Product) Print() {
	fmt.Println(strings.Repeat("=", 20))
	for _, product := range products {
		fmt.Println("Name :", product.Name)
		fmt.Println("Stock :", product.Stock)
		fmt.Println("Price :", product.Price)
		fmt.Println()
	}
}
func main() {
	users := []Repository{
		NewUser("Hacktiv8", "admin@hacktiv8.com", "Jakarta"),
		NewUser("Koinworks", "admin@coinworks.com", "Jakarta"),
	}

	for _, user := range users {
		user.Save()
	}

	products := []Repository{
		NewProduct("Baju", 10, 20000),
		NewProduct("Celana", 21, 50000),
	}
	for _, product := range products {
		product.Save()
	}

	if len(users) > 0 {
		users[0].Print()
	}

	if len(products) > 0 {
		products[0].Print()
	}
}
