package main

import (
	"fmt"
	"sync"
)

var sharedResources = make(map[string]string)

func main8() {
	var wg sync.WaitGroup

	wg.Add(1)

	mu := sync.Mutex{}
	cha := sync.NewCond(&mu)

	go func() {
		defer wg.Done()
		cha.L.Lock()
		for len(sharedResources) < 1 {
			cha.Wait()
			// time.Sleep(1 * time.Millisecond)
		}
		fmt.Println(sharedResources["rc1"])
		cha.L.Unlock()

	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		cha.L.Lock()
		if len(sharedResources) < 2 {
			cha.Wait()
		}
		fmt.Println(sharedResources["rc2"])
		cha.L.Unlock()

	}()
	cha.L.Lock()
	sharedResources["rc1"] = "foo"
	sharedResources["rc2"] = "bar"
	cha.Broadcast()
	cha.L.Unlock()
	wg.Wait()

}
