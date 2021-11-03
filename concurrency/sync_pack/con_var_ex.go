package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRes1 = make(map[string]interface{})

func main5() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		for len(sharedRes1) == 0 {
			time.Sleep(1 * time.Millisecond)
		}
		fmt.Println(sharedRes1["rc1"])
	}()
}
