package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Main process started")
	fruits := []string{"apel", "mangga", "pisang", "jambu"}
	var wg sync.WaitGroup

	wg.Add(len(fruits))
	for i, fruit := range fruits {
		go test(i, fruit, &wg)
	}
	fmt.Println("Number of goroutine :", runtime.NumGoroutine())
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Main process ended")
		wg.Done()
	}(&wg)
	go func() {
		fmt.Println("Main process ended 2")
		wg.Done()
	}()
	wg.Wait()
}

func test(index int, fruit string, wg *sync.WaitGroup) {
	print(index, fruit)
	wg.Done()
}

func print(index int, fruit string) {
	fmt.Printf("index => %d, buah =>%s \n", index, fruit)
}
