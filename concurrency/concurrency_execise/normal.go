package main

import "fmt"

func dowork(input string) {
	for counter := 0; counter < 3; counter++ {
		fmt.Println("Counter", counter, "Input", input)
	}
}

func main() {

	dowork("direct call")
}
