package main

import "fmt"

func gen(ch chan int) {

	for i := 2; ; i++ {
		ch <- i
	}

}

func filter(ins <-chan int, out chan<- int, prime int) {
	for {
		i := <-ins
		if i%prime != 0 {
			out <- i
		}

	}

}

func main() {

	ch := make(chan int)

	go gen(ch)
	x := 6
	for x > 0 {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
		x--
	}

	s := "st"
	for i, v := range s {
		fmt.Printf("%T %T", s[i], v)
	}

}
