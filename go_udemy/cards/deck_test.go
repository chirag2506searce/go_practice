package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected Length: 52\nGot: %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected: Ace of Clubs\nGot: %v", d[0])
	}
}

func TestSaveFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	d := newDeck()
	d.saveFile("_deckTesting")
	loaded := newDeckFromFile("_deckTesting")

	if len(loaded) != 52 {
		t.Errorf("Expected Length: 52\nGot: %v", len(d))
	}
	os.Remove("_deckTesting")
}
