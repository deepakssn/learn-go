package main

func main() {
	cards := newDeck()
	// hand, remaining := deal(cards, 6)
	// cards.print()
	// hand.print()
	// remaining.print()
	cards.writeToFile("mycards")
	newCards := newDeckFromFile("mycards")
	newCards.print()

}
