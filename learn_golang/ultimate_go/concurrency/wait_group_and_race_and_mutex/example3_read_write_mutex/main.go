package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutx sync.RWMutex

var data []string

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {

			writer(i)
		}
		wg.Done()
	}()
	for j := 0; j < 8; j++ {
		go func(id int) {
			reader(id)
		}(j)
	}

	wg.Wait()

}

func reader(id int) {
	rwMutx.Lock()
	fmt.Println("Read", id, len(data))

	rwMutx.Unlock()

}

func writer(i int) {
	rwMutx.Lock()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Write", i)
	data = append(data, fmt.Sprintf("%d", i))
	rwMutx.Unlock()

}
