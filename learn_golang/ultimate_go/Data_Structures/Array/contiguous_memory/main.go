package main

import "fmt"

func main() {

	var five [5]string

	five[0] = "A"
	five[1] = "B"
	five[2] = "C"
	five[3] = "D"
	five[4] = "E"

	for i, v := range five {
		fmt.Printf("value : %s,Address %p Index Add %p\n", v, &v, &five[i])
	}
}
