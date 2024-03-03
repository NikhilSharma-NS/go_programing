package main

type user struct {
	name  string
	email string
}

type admin struct {
	*user // pointer semantic embedding
	level string
}
