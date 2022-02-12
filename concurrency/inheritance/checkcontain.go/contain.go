package main

import (
	"fmt"
	"strings"
)

func main() {

	string1 := "Welcome to world"

	out := strings.Contains(string1, "to")
	fmt.Println(out)
}
