package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var value int
	wg.Add(1)
	go Taskcall(&wg, value)
	wg.Wait()
	fmt.Println("Task Done")

}

func Taskcall(wg *sync.WaitGroup, input int) {
	defer wg.Done()
	input++
	fmt.Println("Task is Running", input)
}
