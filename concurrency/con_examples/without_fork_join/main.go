package main

import (
	"fmt"
	"time"
)

func main() {
	go work()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done waiting")
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("work")
}
