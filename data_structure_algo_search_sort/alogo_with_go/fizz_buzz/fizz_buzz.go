package fizz_buzz

import "fmt"

func Fizz_buzz(n int) {

	for i := 1; i <= n; i++ {
		fmt.Print("Number: ", i)
		if i%3 == 0 {
			fmt.Print("\tFizz")
		}
		if i%5 == 0 {
			fmt.Print("\tBuzz")
		}
		fmt.Println()
	}
}

func Fizz_buzz_with_switch(n int) {

	for i := 1; i <= n; i++ {
		fmt.Print("Number: ", i)
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Print("\tFizz Buzz")
		case i%3 == 0:
			fmt.Print("\tFizz")
		case i%5 == 0:
			fmt.Print("\tBuzz")

		}
		fmt.Println()

	}
}
