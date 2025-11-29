package pattern

import "fmt"

func Pattern5(number int) {
	for i := number; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
