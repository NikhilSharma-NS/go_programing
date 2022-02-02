package main

import (
	"fmt"
	"time"
)

func main() {
	go my()
	time.Sleep(1 * time.Second)
	fmt.Println("finished main")
	fmt.Println(my1())
}

func my() {
	fmt.Println("my")
}

func my1() (a, b string) {
	return "first", "second"
}
