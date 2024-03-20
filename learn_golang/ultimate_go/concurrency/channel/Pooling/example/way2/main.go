package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {

	ch := make(chan string)

	g := runtime.GOMAXPROCS(1)

	for i := 0; i < g; i++ {
		go dowork(i, ch)
	}

	work := 100

	for w := 1; w <= work; w++ {
		ch <- strconv.Itoa(w) + ":work"
		fmt.Println("parent : sent signal :", w)

	}

}

func dowork(i int, ch chan string) {
	for {
		c, w := <-ch
		fmt.Printf("Received By[%d] singal is [%s] \n", i, c)
		if !w {
			break
		}
	}
}
