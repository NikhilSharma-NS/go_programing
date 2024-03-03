package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Println("Hello", u.name, u.email)
}

type admin struct {
	*user
	level string
}

func main() {

	ad := admin{user: &user{"nik", "mail@gmail.com"}}

	ad.user.notify()
	ad.notify()
}
