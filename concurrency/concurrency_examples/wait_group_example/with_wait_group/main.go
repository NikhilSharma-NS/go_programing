package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg1 *sync.WaitGroup, counter int) {

			work(wg1, counter)
		}(&wg, i)
	}
	wg.Wait()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main done")
}

func work(wg *sync.WaitGroup, counter int) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task done", counter)
}
