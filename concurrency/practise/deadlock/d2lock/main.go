package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var m sync.Mutex

	done := make(chan bool)

	fmt.Println("start")
	go func() {
		defer m.Unlock()
		m.Lock()
	}()

	go func() {
		time.Sleep(1)
		m.Lock()
		defer m.Unlock()
		fmt.Println("sinnal")
		done <- true
	}()

	<-done
}

// func main() {

// 	var m sync.Mutex

// 	done := make(chan bool)

// 	fmt.Println("start")
// 	go func() {
// 		m.Lock()
// 	}()

// 	go func() {
// 		time.Sleep(1)
// 		m.Lock()
// 		defer m.Unlock()
// 		fmt.Println("sinnal")
// 		done <- true
// 	}()

// 	<-done
// }
