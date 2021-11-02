package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex

	deposit := func(amount int) {
		mu.Lock()
		balance += amount
		mu.Unlock()
	}

	withdraw := func(amount int) {
		mu.Lock()
		defer mu.Unlock()
		if balance >= amount {
			balance -= amount
		} else {
			fmt.Println("insuicient balance")
		}
	}

	// make 100 deposit of 1
	// and 100  withdraw concurrently
	// run the program and check the result

	wg.Add(100)
	for counter := 0; counter < 100; counter++ {
		go func() {
			defer wg.Done()
			deposit(2)
		}()
	}

	wg.Add(99)
	for counter := 0; counter < 99; counter++ {
		go func() {
			defer wg.Done()
			withdraw(1)
		}()
	}

	wg.Wait()
	fmt.Println("Balance", balance)

}
