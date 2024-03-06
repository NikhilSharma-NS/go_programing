package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {

	g := runtime.GOMAXPROCS(10)
	fmt.Println(g)
}

func main() {
	s := time.Now()
	//fmt.Println(g)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		lowercase()
		wg.Done()
	}()

	go func() {
		uppercase()
		wg.Done()
	}()

	fmt.Println("waiting finish")
	wg.Wait()
	fmt.Println("Main done")

	diff := time.Now().Sub(s)
	fmt.Println("Diff", diff)
}

func lowercase() {
	for count := 0; count < 3; count++ {
		for r := 'a'; r <= 'z'; r++ {
			fmt.Printf("%c\t", r)
		}

	}
}
func uppercase() {
	for count := 0; count < 3; count++ {
		for r := 'A'; r <= 'Z'; r++ {
			fmt.Printf("%c\t", r)
		}

	}
}
