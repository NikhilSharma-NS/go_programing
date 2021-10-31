package main

import "fmt"

func main2() {
	channel := make(chan int)

	go func() {

		for counter := 0; counter < 6; counter++ {
			channel <- counter
		}
		close(channel)
	}()

	for value := range channel {
		fmt.Printf("Value in main thread: %v \n", value)

	}

}
