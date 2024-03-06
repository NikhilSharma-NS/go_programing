package main

import (
	"fmt"
	"os"
	"sync"
)

type Ben struct {
	name string
}

type People interface {
	Speak() bool
}

func (b *Ben) Speak() bool {
	if b.name != "Ben" {
		fmt.Println("Ben say Hello from ", b.name)
		return false
	}
	return true
}

type Jerry struct {
	name string
}

func (b *Jerry) Speak() bool {
	if b.name != "Jerry" {
		fmt.Println("Jerry say Hello from ", b.name)
		return false
	}
	return true
}

func main() {

	ben := Ben{"Ben"}
	jerry := Jerry{"Jerry"}

	pe := People(&ben)

	go func() {
		for {

			pe = &ben

			if !pe.Speak() {
				os.Exit(1)
			}
		}
	}()

	go func() {
		for {

			pe = &jerry

			if !pe.Speak() {
				os.Exit(1)
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Add(1)
	wg.Wait()

}
