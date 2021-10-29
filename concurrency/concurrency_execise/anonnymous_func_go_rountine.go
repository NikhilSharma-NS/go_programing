package main

import (
	"fmt"
	"time"
)

func doWork(input string) {
	for counter := 0; counter < 3; counter++ {
		fmt.Println("Counter", counter, "Input", input)
	}
}

func main() {

	go func() {
		doWork("anonnymous")
	}()

	time.Sleep(1 * time.Microsecond)
	fmt.Println("Done")

}
