package any_base_to_decimal

import "fmt"

func BasetoDecimal(s string, base int) int {
	var sum int
	mul := 1
	for i := len(s) - 1; i >= 0; i-- {
		var val int
		fmt.Sscanf(string(s[i]), "%X", &val)
		sum += mul * val
		mul *= base

	}
	return sum
}
