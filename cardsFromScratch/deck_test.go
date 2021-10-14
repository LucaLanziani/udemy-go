package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLength := 20
	if len(d) != expectedLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedLength, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("First card is not Ace of Spades but is %v", d[0])
	}

	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Last card shuold be a Five of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 20 {
		t.Errorf("Expected 20 cards, got %v", len(loadedDeck))
	}

	os.Remove(filename)
}
