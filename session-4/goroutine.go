package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("main execution started")
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		secondProcess(8)
		wg.Done()
	}()
	go func() {
		firstProcess(8)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Number of goroutine :", runtime.NumGoroutine())
	fmt.Println("main execution ended")
}

func firstProcess(index int) {
	fmt.Println("First process started")

	for i := 0; i < index; i++ {
		fmt.Println("First i=", i)
	}
	fmt.Println("First process ended")
}

func secondProcess(index int) {
	fmt.Println("Second process started")
	for i := 0; i < index; i++ {
		fmt.Println("Second i=", i)
	}
	fmt.Println("Second process ended")
}
