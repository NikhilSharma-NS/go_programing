package main

import "fmt"

// type Animal struct {
// }

// func (a Animal) Eat(input string) {
// 	fmt.Println("Eat", input)
// }

// func (a Animal) Run(input string) {
// 	fmt.Println("run", input)

// }

// type Dog struct {
// 	Animal
// }

// type frok struct {
// 	Animal
// }

// func main() {
// 	d := Dog{}
// 	d1 := frok{}

// 	d.Eat("dog")
// 	d1.Eat("frog")

// }
func main() {
	ch := make(chan int)
	ch <- 7
	val := <-ch
	fmt.Println(val)
}
