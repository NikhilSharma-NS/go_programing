package pattern

import "fmt"

func Pattern2(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
