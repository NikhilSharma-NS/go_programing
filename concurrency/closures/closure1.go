package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	inc := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Printf("Value is %v \n", i)
		}()

	}
	inc(&wg)
	wg.Wait()
	fmt.Print("Done")
}
