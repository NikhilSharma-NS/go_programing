package main

import "fmt"

func main4() {
	// create ch1 and ch2
	// spin gorountine genMsg and replayMsg

	// receive msg on ch2
	ch1 := make(chan int)
	ch2 := make(chan int)
	go genMsg(ch1)
	go receiveMsg(ch1, ch2)
	fmt.Println("Received ", <-ch2)

}
func genMsg(channel chan<- int) {
	// send msg on ch1
	channel <- 1

}

func receiveMsg(channel1 <-chan int, channel2 chan<- int) {
	// receive msg on ch1
	msg := <-channel1
	// sent it on ch2
	channel2 <- msg

}
