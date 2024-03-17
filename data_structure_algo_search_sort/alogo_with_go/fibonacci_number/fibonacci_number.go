package fibonacci_number

import "fmt"

func Fibonacci_number_print(input int) {
	pre := 1
	next := 0
	for i := 0; i < input; {
		if i == 0 {
			fmt.Print(next, "\t")
			i++
		} else if i == 1 {
			fmt.Print(pre, "\t")
			i++
		} else {
			fn := pre + next
			fmt.Print(fn, "\t")
			next = pre
			pre = fn
			i++
		}

	}
}
