package main

import (
	"fmt"
	"runtime"
)

func main() {

	num := 2000

	buf_channel := make(chan string, num)

	g := runtime.GOMAXPROCS(0)
	sem := make(chan bool, g)

	for i := 0; i < num; i++ {
		go func(value int) {
			sem <- true
			{
				buf_channel <- "data"
				fmt.Println("Sent Data", value)

			}
			<-sem

		}(i)

	}

	for num > 0 {
		d := <-buf_channel
		fmt.Println("Received", d, num)
		num--
	}

}
