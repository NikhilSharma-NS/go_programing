package main

import "fmt"

// reader...
type reader interface {
	read(v int)
}

type file struct {
	name string
}

type anotherfile struct {
	name string
}

func (f file) read(v int) {
	fmt.Printf("Name::[%v]  :: marks [%v]\n", v, f.name)
}

func (a anotherfile) read(v int) {
	fmt.Printf("Name is::[%v]  :: marks [%v]\n", v, a.name)
}

func check(r reader) {

	fmt.Println("check called ")
	r.read(10)

}
func main() {
	f := file{
		name: "nikhil",
	}
	f1 := anotherfile{
		name: "jay",
	}

	// f.read(10)
	// f1.read(10)

	check(f)
	check(f1)
}
