package main

import "fmt"

func main() {

	b := make([]byte, 10)
	fmt.Println(b)
	b1 := make([]byte, 2, 3)
	fmt.Println(b1)

	mapd := make(map[string]int)
	mapd["key"] = 0
	fmt.Println(mapd)

}
