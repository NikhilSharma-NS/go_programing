package main

import "fmt"

type printer interface {
	print()
}

type canon struct {
	name string
}

func (c canon) print() {
	fmt.Println(c.name)
}

type epson struct {
	name string
}

func (e *epson) print() {
	fmt.Println(e.name)
}

func main() {

	c := canon{"canon printer"}
	e := epson{"eposon printer"}

	printers := []printer{c, &e}
	c.name = "canon new printer "
	e.name = "eposon new printer"
	for _, p := range printers {
		p.print()
	}
}
