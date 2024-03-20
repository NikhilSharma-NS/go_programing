package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	start := time.Now()
	d := time.Second * 2
	go dowork(d, &wg)
	go dowork(d, &wg)
	wg.Wait()
	fmt.Println(time.Since(start))

}

func dowork(d time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Work Started")
	time.Sleep(d)
	fmt.Println("Work Completed")

}
