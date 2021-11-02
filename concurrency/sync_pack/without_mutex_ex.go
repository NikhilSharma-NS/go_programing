package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main1() {
	runtime.GOMAXPROCS(4)
	var balance int
	var wg sync.WaitGroup
	deposit := func(amount int) {
		balance += amount
	}

	withdraw := func(amount int) {
		balance -= amount
	}

	// make 100 deposit of 1
	// and 100  withdraw concurrently
	// run the program and check the result

	wg.Add(100)
	for counter := 0; counter < 100; counter++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)
	for counter := 0; counter < 100; counter++ {
		go func() {
			defer wg.Done()
			withdraw(1)
		}()
	}

	wg.Wait()
	fmt.Println("Balance", balance)

}
