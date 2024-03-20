package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m1, m2 sync.Mutex

	done := make(chan bool)

	fmt.Println("start")

	go func() {
		m1.Lock()
		defer m1.Unlock()
		time.Sleep(1)
		m2.Lock()
		defer m2.Unlock()

		fmt.Println("signal")

		done <- true
	}()

	go func() {
		m1.Lock()
		defer m1.Unlock()
		time.Sleep(1)
		m2.Lock()
		defer m2.Unlock()

		fmt.Println("signal")

		done <- true
	}()

	<-done
	fmt.Println("done")
	<-done
	fmt.Println("done")
}

// func main() {
// 	var m1, m2 sync.Mutex

// 	done := make(chan bool)

// 	fmt.Println("start")

// 	go func() {
// 		m1.Lock()
// 		defer m1.Unlock()
// 		time.Sleep(1)
// 		m2.Lock()
// 		defer m2.Unlock()

// 		fmt.Println("signal")

// 		done <- true
// 	}()

// 	go func() {
// 		m2.Lock()
// 		defer m2.Unlock()
// 		time.Sleep(1)
// 		m1.Lock()
// 		defer m1.Unlock()

// 		fmt.Println("signal")

// 		done <- true
// 	}()

// 	<-done
// 	fmt.Println("done")
// 	<-done
// 	fmt.Println("done")
// }
