package main

import (
	"fmt"
	"assignment-1/repository"
	"os"
	"strconv"
)

func main(){
	arg := os.Args[1]
	
	if int1, err := strconv.ParseInt(arg, 6, 12); err != nil {
		fmt.Println("Wrong Params")
	} else {
		if int1 > 0 && int(int1) < len(repository.Store) {
			fmt.Println("Name:\t\t", repository.Store[int1 - 1].Name)
			fmt.Println("Address:\t", repository.Store[int1 - 1].Address)
			fmt.Println("Job:\t\t", repository.Store[int1 - 1].Job)
			fmt.Println("Reason:\t\t", repository.Store[int1 - 1].Reason)
		} else {
			fmt.Printf("absen %d not found", int1)
		}
	}
}