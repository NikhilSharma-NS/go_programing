package main

import "fmt"

type data struct {
	name string
	age  int
}

func (d data) diplayNam() {
	fmt.Println(d.name)
}

func (d *data) updateAge(age int) {
	d.age = age
}

func main() {

	d := data{name: "nik"}
	d.updateAge(10)
	d.diplayNam()
	fmt.Println(d.age)

	f1 := d.diplayNam
	f2 := d.updateAge
	f1()
	d.name = "John"
	//	fmt.Println(d.name)
	f1()
	f2(20)
	d.diplayNam()
}
