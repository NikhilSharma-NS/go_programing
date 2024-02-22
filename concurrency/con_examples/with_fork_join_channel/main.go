package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan struct{})
	go func() {
		work()
		done <- struct{}{}
	}()

	<-done
	fmt.Println("wait done")
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("work done")
}
