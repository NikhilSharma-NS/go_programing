package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "ch2"
	}()

	// fmt.Println("received first ", <-ch1)
	// fmt.Println("received second", <-ch2)

	select {
	case m1 := <-ch1:
		fmt.Println("received", m1)
	case m2 := <-ch2:
		fmt.Println("received", m2)

	}

}
