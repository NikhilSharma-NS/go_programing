package main

import (
	"fmt"
	"time"
)

func main() {
	cancelch()
}

func cancelch() {

	ch := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "hello"
		fmt.Println("Send")
	}()

	time_ch := time.After(12 * time.Second)

	select {
	case c := <-ch:
		fmt.Println("Received", c)
	case c_time := <-time_ch:
		fmt.Println("Cancelled", c_time)
	}
}
