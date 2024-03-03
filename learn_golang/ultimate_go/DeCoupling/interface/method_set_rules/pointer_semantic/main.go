package main

import "fmt"

type nofifier interface {
	notify()
}

type User struct {
	name  string
	email string
}

func (u *User) notify() {
	fmt.Println("Name:::" + u.name + "Email::: " + u.email)
}

func sendNotification(n nofifier) {
	n.notify()
}

func main() {

	u := User{"nik", "mail@gmail.com"}

	sendNotification(&u)
}
