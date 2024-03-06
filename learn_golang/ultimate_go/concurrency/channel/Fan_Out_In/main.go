package main

import "fmt"

func main() {

	num := 200000
	buf_channel := make(chan string, num)

	for i := 0; i < num; i++ {
		go func(value int) {
			buf_channel <- "data"
			fmt.Println("Sent Data", value)
		}(i)

	}

	for num > 0 {
		d := <-buf_channel
		fmt.Println("Received", d, num)
		num--
	}

}
