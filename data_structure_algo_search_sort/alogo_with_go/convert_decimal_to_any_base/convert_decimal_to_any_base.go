package convert_decimal_to_any_base

import (
	"fmt"
	"strings"
)

func ConvertDecimalToBase(value int, base int) string {
	var res string
	for value != 0 {

		rem := value % base
		res = fmt.Sprintf("%X%s", rem, res)
		value = value / base
	}

	return res
}

func ConvertDecimalToBase_Constant(value int, base int) string {
	const charset = "0123456789ABCDEF"
	var res string
	for value != 0 {

		rem := value % base
		res = string(charset[rem]) + res

		value = value / base
	}

	return res
}
func ConvertDecimalToBase_Constant_with_String_Builder(value int, base int) string {
	const charset = "0123456789ABCDEF"
	var res strings.Builder
	for value != 0 {
		rem := value % base
		res.WriteByte(byte(charset[rem]))
		value = value / base
	}

	return Reverse(res.String())
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
