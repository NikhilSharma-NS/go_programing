// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	type example struct {
		flag    bool
		counter int16
		pi      float32
	}

	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    false,
		counter: 10,
		pi:      10,
	}

	var e example

	e1 := example{flag: true, counter: 10, pi: 10.0}
	fmt.Println(e, e1, e2)
}
