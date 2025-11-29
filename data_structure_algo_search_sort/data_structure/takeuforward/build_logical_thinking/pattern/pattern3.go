package pattern

import "fmt"

func Pattern3(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
