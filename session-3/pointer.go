package main

import "fmt"

func main() {
	num1 := 10
	fmt.Println("===== BEFORE =====")
	fmt.Println(num1)
	fmt.Println("===== AFTER =====")
	changeNum(&num1, 15)
	fmt.Println(num1)
}

func changeNum(num *int, newNum int) {
	*num = newNum
}
