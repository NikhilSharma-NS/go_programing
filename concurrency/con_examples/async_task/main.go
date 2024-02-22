package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)
	go task1(done)
	go task2(done)
	go task3(done)

	<-done
	<-done
	<-done
	fmt.Println("all waiting done")

}

func task1(done chan bool) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task1")
	done <- true
}
func task2(done chan bool) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task2")
	done <- true
}

func task3(done chan bool) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task3")
	done <- true
}
