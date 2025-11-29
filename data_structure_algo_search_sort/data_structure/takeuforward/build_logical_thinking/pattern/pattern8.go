package pattern

import "fmt"

func Pattern8(number int) {

	for i := number - 1; i >= 0; i-- {
		// left
		for j := 1; j <= number-i-1; j++ {
			fmt.Print(" ")
		}
		// center
		for c := 1; c <= 2*i+1; c++ {
			fmt.Print("*")
		}
		// right
		for k := 1; k <= number-i-1; k++ {
			fmt.Print(" ")
		}
		fmt.Println()

	}

}

// 5-0-1
// 1111*1111
// 111***111
// 11*****11
// 1*******1
// **********
