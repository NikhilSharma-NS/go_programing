package main

import "fmt"

func main() {

	syncCh := make(chan bool)
	done := make(chan bool)
	num := 10
	go even(num, syncCh, done)
	syncCh <- true
	go odd(10, syncCh, done)
	<-done

}

func even(num int, sync chan bool, done chan bool) {

	for i := 1; i <= num; i++ {
		<-sync
		if i%2 == 0 {
			fmt.Println("Number", i, "is even")
		}
		sync <- true
	}

	done <- true

}

func odd(num int, sync chan bool, done chan bool) {
	for i := 1; i <= num; i++ {
		<-sync
		if i%2 != 0 {
			fmt.Println("Number", i, "is odd")
		}
		sync <- true
	}

	done <- true
}
