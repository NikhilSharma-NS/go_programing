package main

import (
	"fmt"
	"runtime"
)

func main() {

	ch := make(chan string)
	v := runtime.GOMAXPROCS(1)
	fmt.Println(v)
	for i := 0; i < v; i++ {
		go func(c int) {
			for i := range ch {
				fmt.Println("Received Signal", i, c)
			}
			fmt.Println("Received shut down Signal", c)
		}(i)
	}

	const work = 100
	for w := 0; w < work; w++ {
		ch <- "data"
		fmt.Println("sent signal", w)
	}
	//close(ch)

}
