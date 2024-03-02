package main

import "fmt"

type user struct {
	name     string
	username string
}

func main() {
	// zero value
	// trying to use this map will result in runtime error(panic)
	var users_zero_value map[string]user
	// construct a map initialized using make
	users_with_make := make(map[string]user)
	// construct a map initialized using empty literal construction
	users_with_empty_literal := map[string]user{}

	//users_zero_value["key"] = user{name: "nik", username: "nikhl"}
	users_with_make["key1"] = user{name: "nik", username: "nikhl"}
	users_with_empty_literal["key2"] = user{name: "nikhil", username: "nikhl"}

	fmt.Println(users_zero_value, users_with_make, users_with_empty_literal)

	users := map[string]user{
		"key1": {name: "jay", username: "virru"},
		"key2": {name: "virru", username: "viru"},
	}

	fmt.Println("us", users)

	user1, exis1 := users["key1"]
	user2, exis2 := users["key3"]

	fmt.Println(user1, exis1)
	fmt.Println(user2, exis2)

	delete(users, "key1")

}
