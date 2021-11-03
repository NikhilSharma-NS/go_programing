package main

import (
	"fmt"
	"sync"
)

var sharedRes = make(map[string]interface{})

func main6() {

	var wg sync.WaitGroup

	mu := sync.Mutex{}
	c := sync.NewCond(&mu)

	wg.Add(1)

	go func() {
		defer wg.Done()
		// suspend goroutine util sharedrsc is populated

		c.L.Lock()

		for len(sharedRes) == 0 {
			//time.Sleep(1 * time.Millisecond)
			c.Wait()
		}
		fmt.Println(sharedRes["rc1"])
		c.L.Unlock()
	}()

	c.L.Lock()
	// writes changes to sharedRSC
	sharedRes["rc1"] = "foo"
	c.Signal()
	c.L.Unlock()

	wg.Wait()
}
