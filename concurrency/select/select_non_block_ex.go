package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		for counter := 0; counter < 3; counter++ {
			time.Sleep(1 * time.Second)
			ch <- "message"
		}
	}()

	// for counter := 0; counter < 3; counter++ {
	// 	m := <-ch
	// 	fmt.Println("msg", m)

	// 	fmt.Println("processing")
	// 	time.Sleep(1500 * time.Millisecond)
	// }
	for counter := 0; counter < 3; counter++ {
		select {
		case m := <-ch:
			fmt.Println("msg", m)
		default:
			fmt.Println("no msg received")
		}
		fmt.Println("processing")
		time.Sleep(1500 * time.Millisecond)

	}

}
