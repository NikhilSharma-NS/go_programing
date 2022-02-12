package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length of deck is 16, but got %v", len(d))
	}
}

func TestDeckPrint(t *testing.T) {
	d := newDeck()
	d.print()
}

func TestSaveToSystemAndReadFromfile(t *testing.T) {
	os.Remove("_deck_testing")
	deck := newDeck()
	deck.saveToSystem("_deck_testing")

	loadedDeck := readFromfile("_deck_testing")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected length of deck is 16, but got %v", len(loadedDeck))
	}

}
