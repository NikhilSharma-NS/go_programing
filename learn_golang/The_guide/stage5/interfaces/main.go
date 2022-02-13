package main

import "fmt"

type englishBot struct {
}

type spanisBot struct {
}

type bot interface {
	getGretting() string
}

func (eb englishBot) getGretting() string {
	return "Hi english"
}

func (sb spanisBot) getGretting() string {
	return "Hello spanish"
}

func printGretting(b bot) {
	fmt.Println(b.getGretting())
}

// func printGretting(eb englishBot) {
// 	fmt.Println(eb.getGretting())
// }

// func printGretting(sb spanisBot) {

// }

func main() {
	eb := englishBot{}
	sb := spanisBot{}
	printGretting(sb)
	printGretting(eb)

}
