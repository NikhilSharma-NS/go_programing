package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	d := time.Second * 2
	dowork(d)
	dowork(d)
	fmt.Println(time.Since(start))
}

func dowork(d time.Duration) {
	fmt.Println("Work Started")
	time.Sleep(d)
	fmt.Println("Work Completed")
}
