package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch := make(chan string)

	go show(ch)
	fmt.Println("parent : recv'd signal :", <-ch)
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")

}

func show(c chan string) {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	c <- "Hi How are you"
	fmt.Println("child : sent signal")

}
