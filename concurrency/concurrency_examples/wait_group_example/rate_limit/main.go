package main

import (
	"fmt"
	"sync"
)

type request func()

func main() {

	requests := map[int]request{}

	for i := 1; i <= 100; i++ {
		f := func(n int) request {
			return func() {
				fmt.Println("request done", n)
			}
		}
		requests[i] = f(i)
	}

	var wg sync.WaitGroup

	max := 10
	for i := 0; i < len(requests); i += max {
		for j := i; j < i+max; j++ {
			wg.Add(1)
			go func(r request) {
				defer wg.Done()
				r()
			}(requests[j+1])
		}
		wg.Wait()
		fmt.Println(max, "request processed")
	}

}
