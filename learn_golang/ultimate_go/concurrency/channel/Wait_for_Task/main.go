package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go func() {
		d := <-ch
		fmt.Println("received:::", d)
	}()

	time.Sleep(100 * time.Microsecond)
	ch <- "task"
	fmt.Println("send signal")

}
