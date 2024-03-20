package main

import "fmt"

func main() {

	ch := make(chan bool)

	go func(flag bool) {
		fmt.Println("start")
		if flag {
			ch <- true
		}

	}(true)

	<-ch
	fmt.Println("done")
}

// ch := make(chan bool)

// go func(flag bool) {
// 	fmt.Println("start")
// 	if flag {
// 		ch <- true
// 	}

// }(false)

// <-ch
// fmt.Println("done")
