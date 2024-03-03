package main

type user struct {
	name  string
	email string
}

type admin struct {
	person user // not embedding
	level  string
}
