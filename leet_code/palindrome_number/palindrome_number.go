package palindrome_number

import (
	"fmt"
	"strconv"
)

func Palindrome_number(num int) bool {
	v := fmt.Sprintf("%v", num)
	fmt.Println("Value is ", v)
	if num < 0 {
		return false
	} else if num < 10 {
		return true
	}

	str := (strconv.Itoa(num))

	for i, j := 0, len(str)-1; i < len(str) && i != j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}

	return true
}
