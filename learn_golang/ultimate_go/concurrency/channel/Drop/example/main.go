package main

import (
	"fmt"
	"strconv"
)

func main() {

	cap := 10
	work := 20

	ch := make(chan string, cap)

	go child(ch)

	for w := 1; w <= work; w++ {
		select {
		case ch <- "work:" + strconv.Itoa(w):
			fmt.Println("sent signal", w)
		default:
			fmt.Println("Droped", w)
		}
	}
	close(ch)
	fmt.Println("parent : sent shutdown signal")

}

func child(ch chan string) {

	for i := range ch {
		fmt.Println("Recived", i)
	}
}
