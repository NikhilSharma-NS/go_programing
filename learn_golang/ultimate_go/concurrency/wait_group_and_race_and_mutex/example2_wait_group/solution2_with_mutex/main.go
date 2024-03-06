package main

import (
	"fmt"
	"sync"
)

func main() {

	var counter int32
	gors := 2

	var wg sync.WaitGroup
	var mutx sync.Mutex

	wg.Add(gors)
	for i := 0; i < gors; i++ {

		go func() {
			for j := 0; j < 2; j++ {
				mutx.Lock()
				counter = counter + 1
				mutx.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
