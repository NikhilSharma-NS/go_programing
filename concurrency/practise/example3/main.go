package main

import "fmt"

func main() {

	ch := make(chan int, 1)

	ch <- 1

	c, flag := <-ch

	fmt.Println(c, flag)

	close(ch)

	c1, flag1 := <-ch

	fmt.Println(c1, flag1)

	// ch := make(chan bool)

	// ch <- true

	// c, flag := <-ch

	// fmt.Println(c, flag)

	// close(ch)

	// c1, flag1 := <-ch

	// fmt.Println(c1, flag1)

}
