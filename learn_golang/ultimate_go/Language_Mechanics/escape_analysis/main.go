package main

import "fmt"

type user struct {
	name  string
	email string
}

func stayonStack() user {
	u := user{
		name:  "nik",
		email: "nik@gmail.com",
	}
	return u
}

func escapeToHeap() *user {
	u := user{
		name:  "nik",
		email: "nik@gmail.com",
	}
	return &u
}

func escapeToHeap_not_Use() *user {
	u := &user{
		name:  "nik",
		email: "nik@gmail.com",
	}
	return u
}

func main() {
	fmt.Println(stayonStack())
	fmt.Println(escapeToHeap())
	fmt.Println(escapeToHeap_not_Use())

	x := stayonStack()
	y := escapeToHeap()
	z := escapeToHeap_not_Use()
	fmt.Println("x,y,z", x, y, z)
	x.email = "jay"
	y.email = "jay"
	z.email = "jay"
	fmt.Println("x,y,z", x, y, z)

}
