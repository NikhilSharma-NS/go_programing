package reverse_a_string

import "strings"

func Reverse_a_string_with_Rune(s string) string {
	var rev string
	for _, v := range s {
		rev = string(v) + rev
	}

	return rev
}

func Reverse_a_string_with_byte(s string) string {
	var rev string
	for i := 0; i < len(s); i++ {
		rev = string(s[i]) + rev
	}
	return rev
}

func Reverse_a_string_with_string_builder(s string) string {
	var sb strings.Builder

	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}
