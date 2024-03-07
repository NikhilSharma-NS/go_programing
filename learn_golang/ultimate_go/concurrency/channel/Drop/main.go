package main

import (
	"fmt"
	"time"
)

func main() {
	drop()
}

func drop() {
	const cap = 5
	ch := make(chan string, 5)

	go func() {
		for i := range ch {
			fmt.Println("Received", i)
		}
	}()
	const work = 20
	for j := 0; j < work; j++ {
		select {
		case ch <- "hello":
			fmt.Println("send,", j)
		default:
			fmt.Println("dropped")
		}
	}

	time.Sleep(time.Second)

}
