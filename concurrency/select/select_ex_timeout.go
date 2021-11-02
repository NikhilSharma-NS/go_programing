package main

import (
	"fmt"
	"time"
)

func main2() {

	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "ch"
	}()

	select {
	case m1 := <-ch:
		fmt.Println("m1", m1)
	case <-time.After(1 * time.Second):
		fmt.Println("TIme out")

	}
}
