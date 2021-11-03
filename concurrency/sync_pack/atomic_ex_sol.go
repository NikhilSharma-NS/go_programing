package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main4() {
	runtime.GOMAXPROCS(4)

	var counter uint64
	var wg sync.WaitGroup

	for couneter1 := 0; couneter1 < 50; couneter1++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("counter", counter)

}
