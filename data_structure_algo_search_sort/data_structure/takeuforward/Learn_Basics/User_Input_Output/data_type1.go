package input_output

import "fmt"

func TakeInput() {
	var i int
	fmt.Print("Input:: ")
	fmt.Scan(&i)
	fmt.Println("Your Number is", i)
}
