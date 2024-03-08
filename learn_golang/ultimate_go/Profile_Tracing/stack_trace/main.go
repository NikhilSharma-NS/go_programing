package main

import "fmt"

func main() {

	work(10)

}

func work(val int) {
	fmt.Println(val)
	panic(0)
}
