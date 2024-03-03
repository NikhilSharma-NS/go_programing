package main

type user struct {
	name  string
	emial string
}
type admin struct {
	user  // value semantic embedding
	level string
}
