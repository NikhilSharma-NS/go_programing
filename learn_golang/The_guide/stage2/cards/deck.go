package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// create a new type of deck
// which is slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Heart", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardV := range cardValues {
		for _, cardSu := range cardSuits {
			cards = append(cards, cardV+" of "+cardSu)
		}
	}
	return cards
}

func (d deck) print() {
	for index, v := range d {
		fmt.Println("deck", index, v)
	}
}

func handSize(handSizei int, d deck) (start deck, end deck) {
	return d[:handSizei], d[handSizei:]
}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")

}

func (d deck) saveToSystem(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromfile(fileName string) deck {
	bytes, err := ioutil.ReadFile(fileName)
	if err == nil {
		return deck(strings.Split(string(bytes), ","))
	}
	return nil
}

func (d deck) shuffle() {
	for index, v := range d {
		newPos := rand.Intn(len(d) - 1)
		d[index], d[newPos] = d[newPos], v
	}
}

func (d deck) shufflewithRandom() {
	for index, v := range d {
		sou := rand.NewSource(time.Now().UnixNano())
		r := rand.New(sou)
		newPos := r.Intn(len(d) - 1)
		d[index], d[newPos] = d[newPos], v
	}
}
