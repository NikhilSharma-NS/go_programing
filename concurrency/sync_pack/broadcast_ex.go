package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedResource = make(map[string]string)

func main7() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		for len(sharedResource) < 1 {
			time.Sleep(1 * time.Millisecond)
		}
		fmt.Println(sharedResource["rc1"])

	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		if len(sharedResource) < 2 {
			time.Sleep(1 * time.Millisecond)

		}
		fmt.Println(sharedResource["rc2"])

	}()

	sharedResource["rc1"] = "foo"
	sharedResource["rc2"] = "bar"

	wg.Wait()

}
