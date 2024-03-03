package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	*user
	levl string
}

func (u *user) notify() {
	fmt.Println("User Notification", u.name, u.email)
}

func (a *admin) notify() {

	fmt.Println("Admin Notification", a.name, a.email)
}

func sendNotification(n notifier) {

	n.notify()
}

func main() {

	adm := &admin{
		user: &user{name: "nik", email: "mail@gmail.com"},
		levl: "level",
	}

	sendNotification(adm)

	adm1 := admin{
		user: &user{name: "nik", email: "mail@gmail.com"},
		levl: "level",
	}

	sendNotification(&adm1)

	adm.user.notify()

}
