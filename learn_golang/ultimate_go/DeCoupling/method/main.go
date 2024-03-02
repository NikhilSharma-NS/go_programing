package main

import "fmt"

type User struct {
	userid int
	name   string
	mail   string
}

func (u User) notify() {
	fmt.Printf("name:[%v], id:[%v], mail[%v]\n", u.name, u.userid, u.mail)
}

func (u *User) chaneMail(email string) {
	u.mail = email
}
func main() {
	// Values of type user
	u := User{userid: 123, name: "nikhil"}
	u.mail = "mail@gmail.com"
	u.chaneMail("changed@gmail.com")
	u.notify()
	// Pointer of type user
	u1 := &User{userid: 123, name: "nikhil"}
	u1.mail = "mail@gmail.com"
	u1.chaneMail("changed@gmail.com")
	u1.notify()
	// Pointer of type user
	u2 := &User{userid: 123, name: "nikhil"}
	u2.mail = "mail@gmail.com"
	u2.chaneMail("changed@gmail.com")
	// through pointer
	(*u2).notify()

}

// Function is called a method when that function has a Receiver
