package main

import "fmt"

func newDeck() []string {
	deck := []string{}
	suits := [4]string{"Clubs", "Diamonds", "Spades", "Hearts"}
	values := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, value+" of "+suit)
		}
	}

	return deck
}

func printDeck(deck []string) {
	for _, card := range deck {
		fmt.Println(card)
	}
}

func main() {
	deck := newDeck()
	printDeck(deck)
}
