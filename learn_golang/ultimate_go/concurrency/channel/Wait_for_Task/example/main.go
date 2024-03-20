package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch := make(chan string)

	go dowork(ch)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "hi how are you"
	fmt.Println("parent : sent signal")
	time.Sleep(time.Second)

}

func dowork(ch chan string) {
	fmt.Println("child : recv'd signal :", <-ch)

}
