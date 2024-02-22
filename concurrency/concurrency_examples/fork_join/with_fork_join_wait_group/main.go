package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		work()
	}()

	time.Sleep(100 * time.Millisecond)
	wg.Wait()
	fmt.Println("waiting done")

}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("work done")
}
