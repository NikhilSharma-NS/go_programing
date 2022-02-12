package main

import (
	"practise/tasks"
	"time"
)

func main() {

	//tasks.StackprintLIFO()
	//fmt.Println("First")
	channel := make(chan bool)
	go tasks.Task1(channel)
	channel <- true

	ch := make(chan int)
	go tasks.Task2(ch)
	ch <- 1

	time.Sleep(1)
}
