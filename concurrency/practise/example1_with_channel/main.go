package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	ch := make(chan string)

	go dowork(time.Second*2, "hello1", ch)
	go dowork(time.Second*2, "hello2", ch)
	result := <-ch
	result1 := <-ch
	fmt.Println(result)
	fmt.Println(result1)
	fmt.Println(time.Since(t))

}

func dowork(d time.Duration, s string, c chan string) {

	fmt.Println("started work")
	time.Sleep(d)
	fmt.Println("done work")
	c <- s

}
