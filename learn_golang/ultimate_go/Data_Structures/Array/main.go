package main

import "fmt"

func main() {
	// slice
	var values []int
	// array
	var dats [10]string

	fmt.Println(values)
	fmt.Println(dats)
	fmt.Println("Length", len(values), "Data", values)
	fmt.Println("Length", len(dats), "Data", dats)

	values = append(values, 10)
	//dats = append(dats, "1")
	fmt.Println("Length", len(values), "Data", values)
	fmt.Println("Length", len(dats), "Data", dats)

}
