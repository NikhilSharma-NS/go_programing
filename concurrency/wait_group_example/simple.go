package main

import (
	"fmt"
)

func main() {

	var data int

	go func() {
		data++
	}()
	//time.Sleep(1 * time.Microsecond)
	if data == 0 {
		//time.Sleep(1 * time.Microsecond)

		fmt.Printf("Data is %v", data)
	}

}
