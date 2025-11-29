package pattern

import "fmt"

func Pattern1(number int) {

	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
