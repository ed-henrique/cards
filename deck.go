package main

import "fmt"

type deck []string

func (d deck) new() deck {
	cards := deck{}
	suits := deck{"Clubs", "Diamonds", "Spades", "Hearts"}
	values := deck{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
