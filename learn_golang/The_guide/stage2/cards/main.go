package main

import "fmt"

func main() {

	cards := newDeck()
	fmt.Println(cards)
	cards = append(cards, newCard())
	fmt.Println(cards)

	for index, v := range cards {
		fmt.Println(index, v)
	}

	// print with deck
	cards.print()
	cardinput := deck{"apple", "banana", "grape", "orange", "mango"}

	fmt.Println(cardinput[:3])
	fmt.Println(cardinput[3:])
	start, end := handSize(1, cardinput)
	start.print()
	end.print()

	stringVa := cardinput.toString()
	fmt.Println(stringVa)
	cardinput.saveToSystem("output.txt")
	cardsfromFile := readFromfile("output.txt")
	fmt.Println(cardsfromFile)
	cardsfromFile.shuffle()
	fmt.Println(cardsfromFile)

	cardsfromFileN := readFromfile("output.txt")
	cardsfromFileN.shufflewithRandom()
	fmt.Println(cardsfromFileN)

}

func newCard() string {
	return "new Card"
}
