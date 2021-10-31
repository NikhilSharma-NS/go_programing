package main

import (
	"fmt"
	"time"
)

func main2() {

	fv := mywork
	go fv("mywork")
	time.Sleep(1 * time.Microsecond)
	fmt.Println("done")
}

func mywork(input string) {
	for counter := 0; counter < 3; counter++ {
		fmt.Println("counter", counter, "Input", input)
	}

}
