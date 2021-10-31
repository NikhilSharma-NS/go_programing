package main

import "fmt"

func main() {
	// create channel owner ,which creates channel
	// return receive only channle to caller
	// spins a goroutine ,which
	// writes data into channel and
	// closes the channel when done

	owner := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for counter := 0; counter < 5; counter++ {
				ch <- counter
			}
		}()
		return ch

	}
	consumer := func(ch <-chan int) {

		// read value of channel
		for value := range ch {
			fmt.Println("Received ", value)
		}
		fmt.Println("Done receiving ")
	}

	ch := owner()
	consumer(ch)

}
