package main

import (
	"fmt"
	"log"
)

type product struct {
	Name  string
	Stock int
	Price int
}

func (p *product) buy(customer string, amount int, c chan string) {
	defer catchError()
	if p.Stock-amount < 0 {
		err := fmt.Sprintf("%s cannt buy %d %s because run out of stock - %d left", customer, amount, p.Name, p.Stock)
		c <- ""
		panic(err)
	}
	newStock := p.Stock - amount
	msg := fmt.Sprintf("%s buy %d in %s - %d left", customer, amount, p.Name, newStock)
	p.Stock = newStock
	c <- msg
}

func main() {
	c := make(chan string, 3)
	baju := &product{Name: "Baju", Stock: 10, Price: 20000}
	go baju.buy("Hacktiv8", 3, c)
	go baju.buy("Koinworks", 5, c)
	go baju.buy("NooBeeID", 7, c)

	for i := 0; i < cap(c); i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Stok sekarang : ", baju.Stock)
}

func catchError() {
	if r := recover(); r != nil {
		log.Fatal("error occured :", r)
	}
}
