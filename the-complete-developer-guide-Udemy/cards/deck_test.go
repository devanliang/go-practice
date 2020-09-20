package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first cards if Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first cards if Ace of Spades, but got %v", d[len(d)-1])
	}
}

func TestSaceToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting") // 刪除所有測試file

	deck := newDeck()
	deck.saveToFile("_decktesting") // 儲存file

	loadedDeck := newDeckFromFile("_decktesting") // loading file

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
