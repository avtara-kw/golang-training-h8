package main

import (
	"fmt"
	"log"
	"errors"
)

func main() {
	defer catchErr()

	type product struct {
		name string
		stock int
		price int
	}

	baju := product{name:"baju", stock:9, price:200}
	c := make(chan map[string]interface{})
	go buy("joko", 3, c)
	go buy("sasa", 6, c)
	go buy("hardi", 3, c)

	fmt.Println(baju)
	for i:=0;i<3;i++{
		temp := <- c
		if baju.stock >= temp["total"].(int) {
			baju.stock -= temp["total"].(int)
			fmt.Println(temp["name"], " beli ", temp["total"]," - sisa ", baju.stock)
		}else{
			panic(errors.New(fmt.Sprintf("habis, sisa %d\n", baju.stock)))
		}
	}
}

func buy(name string, total int, c chan map[string]interface{}) {
	temp := make(map[string]interface{})
	temp["name"] = name
	temp["total"] = total
	c <- temp
}

func catchErr() {
	if r := recover(); r != nil {
		log.Println("Error occured:", r)
	} else {
		log.Println("Application running smooth")
	}
}