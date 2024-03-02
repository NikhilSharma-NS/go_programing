package main

import "fmt"

type user struct {
	like int
}

func main() {

	users := make([]user, 1)

	ptUser0 := &users[0]
	ptUser0.like++

	for i := range users {
		fmt.Printf("Users %d: %v\n", i, users[i])
	}
	fmt.Println("---------------------------------")

	users = append(users, user{})
	// after appen it will have copy
	ptUser0.like++
	for i := range users {
		users[0].like = 5
		fmt.Printf("Users %d: %v \n", i, users[i])
	}
	for i, v := range users {
		fmt.Printf("Users %d: %v\n", i, v)
	}
}
