package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

type Cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}

func (d *Dog) Speak() {
	fmt.Println("Dog::", d.Name, d.IsMammal, d.PackFactor)
}
func (c *Cat) Speak() {
	fmt.Println("Cat::", c.Name, c.IsMammal, c.ClimbFactor)
}

func main() {
	sp := []Speaker{
		&Dog{Name: "d1", IsMammal: true, PackFactor: 1},
		&Cat{Name: "c1", IsMammal: true, ClimbFactor: 1}}

	for _, speaker := range sp {
		speaker.Speak()
	}

}
