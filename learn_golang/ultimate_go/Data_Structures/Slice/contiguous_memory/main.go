package main

import "fmt"

func main() {

	five := make([]string, 5, 8)

	five[0] = "A"
	five[1] = "B"
	five[2] = "C"
	five[3] = "D"
	five[4] = "E"

	for i := range five {
		fmt.Printf("value : %s,Address %p Index  %d\n", five[i], &five[i], i)
	}

}
