package reverse_integer

import "math"

func ReverseInteger(input int) (rev int) {

	for input != 0 {
		mod := input % 10
		if input > 0 {
			if rev > math.MaxInt32/10 {
				return 0
			}
		} else {
			if rev < math.MinInt32 {
				return 0
			}
		}

		rev = mod + rev*10
		input = input / 10

	}
	return rev

}
