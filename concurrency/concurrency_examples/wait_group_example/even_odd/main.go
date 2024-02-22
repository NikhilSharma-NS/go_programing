package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("-------------------------start-------------")
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)

		go func(wg1 *sync.WaitGroup, counter int) {

			if counter%2 == 0 {
				even(wg1, counter)
			} else {
				odd(wg1, counter)
			}
		}(&wg, i)
	}
	wg.Wait()
	fmt.Println("-------------------------end-------------")
}

func even(wg *sync.WaitGroup, counter int) {
	wg.Done()
	fmt.Println("this is even:::::", counter)
}

func odd(wg *sync.WaitGroup, counter int) {
	wg.Done()
	fmt.Println("this is odd::::::", counter)
}
