package main

import (
	"fmt"
	"sync"
	"time"
)

var wg *sync.WaitGroup

func main() {
	t := time.Now()
	ch := make(chan string)
	wg = &sync.WaitGroup{}
	wg.Add(2)
	go dowork(time.Second*2, "hello1", ch)
	go dowork(time.Second*2, "hello2", ch)

	go func() {

		for c := range ch {
			fmt.Println(c)

		}
	}()

	wg.Wait()
	//	close(ch)

	fmt.Println(time.Since(t))

}

func dowork(d time.Duration, s string, c chan string) {
	defer wg.Done()
	fmt.Println("started work")
	time.Sleep(d)
	fmt.Println("done work")
	c <- s

}
