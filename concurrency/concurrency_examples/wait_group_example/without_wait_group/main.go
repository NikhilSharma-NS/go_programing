package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go work(i)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main done")
}

func work(counter int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task done", counter)
}
