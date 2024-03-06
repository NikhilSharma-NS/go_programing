package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go func() {
		time.Sleep(1000 * time.Millisecond)
		ch <- "Hello"
	}()
	fmt.Println("waiting-----------------")
	d := <-ch
	fmt.Println("Received", d)
	fmt.Println("done-----------------")
}
