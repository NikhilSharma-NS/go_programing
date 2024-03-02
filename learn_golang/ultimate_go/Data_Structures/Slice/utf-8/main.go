package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "means world $"
	for i, r := range s {
		fmt.Printf("%2d:%q\n \n", i, r)
	}

	for i, r := range s {
		var buf [utf8.UTFMax]byte
		r1 := utf8.RuneLen(r)
		si := i + r1
		copy(buf[:], s[i:si])
		fmt.Printf("%2d:%q:code:%#6x,%#v\n \n", i, r, r, buf[:r1])
	}
}
