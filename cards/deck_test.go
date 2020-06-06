package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected 16 cards in a new deck, but got only %v", len(d))
	}

}

func TestNewDeckFromFileAndWriteToFile() {

	d := newDeck()
	os.Remove("_testDeck")
	d.writeToFile("_testDeck")
	loadedDeck := newDeckFromFile("_testDeck")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in a new deck, but got only %v", len(loadedDeck))
	}
	os.Remove("_testDeck")

}
