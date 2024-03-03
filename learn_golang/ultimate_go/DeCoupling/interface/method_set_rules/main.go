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
	// cannot use u (variable of type User) as nofifier value in argument to sendNotification: missing method notify (notify has pointer receiver)compiler
	sendNotification(u)
}

// for any value of type T , only those methods implemented with a value receiver for that type belong to th method set to that value
// for any address of type T , all method implemnted for that type belong to the method set of that value
