package main

type IP []byte
type IPMask []byte

func main() {

	i := IP{}
	i.Mask(IPMask{})

}

func (ip IP) Mask(mask IPMask) IP {

	return []byte{}

}

// refernce type and BuildIn(int,bool) type on value semantic
// map channel pointer semantic
// struct type -> accoridn to need we can have value or pointeer
