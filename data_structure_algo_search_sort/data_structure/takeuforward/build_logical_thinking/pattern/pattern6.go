package pattern

import "fmt"

func Pattern6(number int) {
	for i := number; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
