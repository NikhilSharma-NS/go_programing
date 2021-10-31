package main

import (
	"fmt"
	"time"
)

func doWork1(input string) {
	for counter := 0; counter < 3; counter++ {
		fmt.Println("Counter", counter, "Input", input)
	}
}

func main1() {

	go func() {
		doWork1("anonnymous")
	}()

	time.Sleep(1 * time.Microsecond)
	fmt.Println("Done")

}
