package main

import (
	"fmt"
	"sync"
)

var score = make(map[string]int)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {

		for i := 0; i < 1000; i++ {
			score["A"]++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			score["B"]++
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(score)

}
