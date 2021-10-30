package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go func(in int) {
			defer wg.Done()
			fmt.Println(in)
		}(i)
	}
	wg.Wait()

}
