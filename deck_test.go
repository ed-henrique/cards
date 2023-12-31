package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected 'Ace of Clubs' as the first card, but got %v", d[0])
	}
	if d[len(d)-1] != "King of Hearts" {
		t.Errorf("Expected 'King of Hearts' as the last card, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")
	ld := newDeckFromFile("_decktesting")

	if len(ld) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(ld))
	}

	os.Remove("_decktesting")
}
