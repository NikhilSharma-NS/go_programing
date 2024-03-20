package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	n := 1000

	ch := make(chan string, n)

	for i := 1; i <= n; i++ {
		go dowork(i, ch)
	}
	for n > 0 {
		fmt.Println("parent : recv'd signal :", <-ch)

		n--
	}
	close(ch)
}

func dowork(v int, c chan string) {

	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	c <- strconv.Itoa(v) + ":Data"
	fmt.Println("child : sent signal :", v)

}
