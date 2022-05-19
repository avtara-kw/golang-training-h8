package main

import "fmt"

var status = 0
var total = 100

func main() {
	c := make(chan string)
	stat := ""
	for {
		go inquire(c, 11)

		status := <-c
		fmt.Println(status)
		if status == "DECLINED" {
			stat = "DECLINED"
			break
		} else if status == "ACCEPTED" {
			stat = "ACCEPTED"
			break
		}
	}
	fmt.Println(stat)

	close(c)
}

func inquire(c chan string, amount int) {
	total -= amount
	fmt.Println(total)
	if total == 0 {
		c <- "ACCEPTED"
	} else if total < 0 {
		c <- "DECLINED"
	}
	c <- "PENDING"
}
