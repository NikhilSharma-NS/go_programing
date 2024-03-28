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

func Palindrome_number1(num int) bool {
	value := num
	if num < 0 {
		return false
	}
	rev := 0
	for num != 0 {
		mod := num % 10
		rev = mod + rev*10
		num = num / 10
	}
	if rev == value {
		return true
	} else {
		return false
	}
}
