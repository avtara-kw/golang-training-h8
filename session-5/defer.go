package main

import (
	"fmt"
	"os"
)

func main() {
	defer deferFunction()
	fmt.Println("Hello")
	os.Exit(1)
}

func deferFunction() {
	defer print()
	fmt.Println("from defer function")
	return
}

func print() {
	fmt.Println("Call from defer")
}
