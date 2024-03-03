package main

import "fmt"

type notifier interface {
	notify()
}

func sendNotification(n notifier) {
	n.notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Println("User Notification", u.name, u.email)
}

type admin struct {
	*user
	level string
}

func main() {

	ad := admin{user: &user{"nik", "mail@gmail.com"}}

	sendNotification(ad)
	sendNotification(&ad)
}
