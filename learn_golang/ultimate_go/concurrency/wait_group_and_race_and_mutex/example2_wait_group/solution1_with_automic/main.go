package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int32
	gors := 2

	var wg sync.WaitGroup
	wg.Add(gors)
	for i := 0; i < gors; i++ {

		go func() {
			for j := 0; j < 2; j++ {
				atomic.AddInt32(&counter, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
