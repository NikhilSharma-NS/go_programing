package main

import (
	"fmt"
	"time"
)

//waitforfinished
func main() {

	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Second)
		fmt.Println("doing employee")
		close(ch)
		fmt.Println("employee sent signal ")
	}()

	_, wd := <-ch

	fmt.Println("manager received", wd)
	//	time.Sleep(time.Second)

}
