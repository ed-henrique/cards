package main

func main() {
	cards := newDeck()
	cards.print()

	hand, cards := deal(cards, 5)

	hand.print()
	cards.print()
}
