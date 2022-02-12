package tasks

import "fmt"

func Task1(ch chan bool) {
	fmt.Println(<-ch)
}

func Task2(chans chan int) {
	if <-chans == 1 {
		fmt.Println("Channel Received")
	}
}
