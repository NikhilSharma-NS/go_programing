package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)
	go func(a, b int) {
		c := a + b
		channel <- c
		fmt.Println("SUm", c)
	}(1, 2)
	fmt.Printf("Main function is having SUM %v", <-channel)
	time.Sleep(1 * time.Millisecond)
}
