package main

func main() {
	cards := newDeck()
	// cards.print()
	hand, cards := deal(cards, 5)
	// hand.print()
	// cards.print()
	hand.saveToFile("test")
	//cards2 := newDeckFromFile("test")
	//cards2.print()
	cards.shuffle()
	cards.print()
}
