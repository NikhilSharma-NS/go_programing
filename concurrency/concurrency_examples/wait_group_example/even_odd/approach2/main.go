package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	rec := make(chan int)
	done := make(chan bool)
	go receive_data(arr, rec, done)
	go processAndCheckData(rec)
	<-done

}

func receive_data(arr []int, receiver chan int, done chan bool) {

	for i := 0; i < len(arr); i++ {
		receiver <- arr[i]
	}
	done <- true
}

func processAndCheckData(receiver chan int) {

	for data := range receiver {
		if data%2 == 0 {
			fmt.Println("Number is even", data)
		} else {
			fmt.Println("Number is odd", data)
		}
	}
	close(receiver)

}
