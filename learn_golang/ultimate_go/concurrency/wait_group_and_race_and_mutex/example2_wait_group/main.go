package main

import (
	"fmt"
	"sync"
)

func main() {

	var counter int
	gors := 2

	var wg sync.WaitGroup
	wg.Add(gors)
	for i := 0; i < gors; i++ {

		go func() {
			for j := 0; j < 2; j++ {
				counter = counter + 1

			}
			wg.Done()

		}()

	}
	wg.Wait()
	fmt.Println(counter)
}
