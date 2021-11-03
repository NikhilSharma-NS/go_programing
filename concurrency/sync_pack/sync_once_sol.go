package main

import (
	"fmt"
	"sync"
)

func main10() {
	var wg sync.WaitGroup
	var synco sync.Once

	load := func() {
		fmt.Println("Run Once")
	}
	loadmul := func() {
		fmt.Println("Run multiple")
	}

	for counter := 0; counter < 10; counter++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			synco.Do(load)
			loadmul()
		}()
	}
	wg.Wait()

}
