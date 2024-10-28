package test

import (
	"github.com/LilMonk/go-tut/pkg/probability"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	deck := probability.CreateDeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(deck))
	}

	suitCount := make(map[string]int)
	rankCount := make(map[string]int)

	for _, card := range deck {
		suitCount[card.Suit]++
		rankCount[card.Rank]++
	}

	for _, suit := range probability.SUITS {
		if suitCount[string(suit)] != 13 {
			t.Errorf("Expected 13 cards of suit %s, but got %d", string(suit), suitCount[string(suit)])
		}
	}

	for _, rank := range probability.RANKS {
		if rankCount[string(rank)] != 4 {
			t.Errorf("Expected 4 cards of rank %s, but got %d", string(rank), rankCount[string(rank)])
		}
	}
}

func TestCreateHandsCombinations(t *testing.T) {
	deck := probability.CreateDeck()
	handSize := 5
	hands := probability.CreateHandsCombinations(deck, handSize)

	expectedCombinations := 2598960 // 52 choose 5
	if len(hands) != expectedCombinations {
		t.Errorf("Expected %d combinations, but got %d", expectedCombinations, len(hands))
	}

	for _, hand := range hands {
		if len(hand) != handSize {
			t.Errorf("Expected hand size of %d, but got %d", handSize, len(hand))
		}
		cardMap := make(map[string]bool)
		for _, card := range hand {
			cardStr := card.String()
			if cardMap[cardStr] {
				t.Errorf("Duplicate card found in hand: %s", cardStr)
			}
			cardMap[cardStr] = true
		}
	}
}
