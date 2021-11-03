package main

import (
	"fmt"
	"sync"
)

func main9() {
	var wg sync.WaitGroup

	load := func() {
		fmt.Println("Run Once")
	}

	for counter := 0; counter < 10; counter++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			load()
		}()
	}
	wg.Wait()

}
