package main

import "fmt"

func main() {

	// color := map[string]string{
	// 	"key": "value",
	// 	"ke1": "va",
	// }

	// fmt.Println(color)

	// var colors map[string]string
	// fmt.Println(colors)

	colors := make(map[int]string)
	colors[1] = "va"

	delete(colors, 3)
	fmt.Println(colors)

	for key, v := range colors {
		fmt.Println(key, "-", v)
	}

}
