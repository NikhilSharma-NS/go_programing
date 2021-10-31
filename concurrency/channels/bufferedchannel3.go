package main

import "fmt"

func main3() {
	channel := make(chan int, 6)

	go func() {
		defer close(channel)
		for counter := 1; counter <= 6; counter++ {
			fmt.Println("Sending ", counter)
			channel <- counter
		}
	}()

	for v := range channel {
		fmt.Printf("Received %v\n", v)

	}

}
