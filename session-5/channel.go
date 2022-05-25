package main

import (
	"fmt"
	"time"
)

func main() {
	buffer()
}

func unbuffer() {
	c1 := make(chan int)
	go func(c chan int) {
		fmt.Println("Goroutine started sending data to channel")
		c1 <- 10
		fmt.Println("Goroutine after sending data to channel")
	}(c1)
	fmt.Println("Main goroutine sleep for 2 second")
	time.Sleep(2 * time.Second)
	fmt.Println("Main Goroutine started receiving data from channel")
	msg := <-c1
	fmt.Printf("Main Goroutine after receiving data from channel with data :%d\n", msg)

	close(c1)
	time.Sleep(time.Second)
}

func buffer() {
	c1 := make(chan int, 4)
	go func(c chan int) {
		for i := 0; i < 5; i++ {
			fmt.Printf("func goroutine %d started sending data into channel\n", i)
			c1 <- i
			fmt.Printf("func goroutine %d after sending data into channel\n", i)
		}
		close(c)
	}(c1)
	fmt.Println("Main goroutine sleep for 2 second")
	time.Sleep(2 * time.Second)
	fmt.Println("Main Goroutine started receiving data from channel")
	for v := range c1 {
		fmt.Printf("main goroutine received value from channel:%d\n", v)
	}
}

func print(c <-chan string) {
	fmt.Println(<-c)
}

func introduce(student string, c chan<- string) {
	result := fmt.Sprintf("Hello, my name is %s", student)

	c <- result
}
