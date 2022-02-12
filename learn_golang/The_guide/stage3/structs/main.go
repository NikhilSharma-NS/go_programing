package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
	values []string
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%v", p)
}

func (p1 *person) updateName(newfirstname string) {
	p1.firstName = newfirstname
}

func main() {

	// p := person{firstName: "nikhi", lastName: "sharma"}
	// p.firstName = "jay"
	// p1 := person{"nikhi", "sharma"}
	// p2 := person{}
	// fmt.Println(p)
	// fmt.Println(p1)
	// fmt.Println(p2)

	p := person{
		firstName: "Nikhi",
		lastName:  "Sharma",
		contactInfo: contactInfo{
			email:   "nk@gmail.com",
			zipCode: 411057,
		},
		values: []string{"my"},
	}
	fmt.Println(p)
	p2 := &p
	p.updateName("nicky")
	p.print()
	fmt.Println()
	p2.print()
	fmt.Println()

}
