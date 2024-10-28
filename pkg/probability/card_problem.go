// Package probability provides utilities for working with probabilities and combinations
// in card games.
//
// This package includes functionality to create a deck of cards, represent individual cards,
// and generate all possible combinations of hands from a deck of cards.

package probability

// SUITS represents the four suits in a standard deck of cards.
const SUITS = "♠♥♦♣"

// RANKS represents the thirteen ranks in a standard deck of cards.
const RANKS = "A23456789TJQK"

// Card represents a playing card with a suit and a rank.
type Card struct {
	Suit string // Suit of the card (e.g., ♠, ♥, ♦, ♣)
	Rank string // Rank of the card (e.g., A, 2, 3, ..., K)
}

// NewCard creates a new Card with the specified suit and rank.
func (c Card) NewCard(suit string, rank string) Card {
	return Card{Suit: suit, Rank: rank}
}

// String returns the string representation of the card in the format "RankSuit".
func (c Card) String() string {
	return c.Rank + c.Suit
}

// CreateDeck creates a standard 52-card deck.
func CreateDeck() []Card {
	deck := make([]Card, 0, 52)
	for _, s := range SUITS {
		for _, r := range RANKS {
			deck = append(deck, Card{Suit: string(s), Rank: string(r)})
		}
	}
	return deck
}

// CreateHandsCombinations generates all possible combinations of hands of a given size
// from the provided deck of cards.
func CreateHandsCombinations(deck []Card, handSize int) [][]Card {
	var result [][]Card
	var current []Card
	createHandsCombinationsHelper(deck, handSize, 0, current, &result)
	return result
}

// createHandsCombinationsHelper is a recursive helper function that generates combinations
// of hands from the deck of cards.
func createHandsCombinationsHelper(deck []Card, handSize int, start int, current []Card, result *[][]Card) {
	if len(current) == handSize {
		*result = append(*result, append([]Card{}, current...))
		return
	}
	for i := start; i < len(deck); i++ {
		current = append(current, deck[i])
		createHandsCombinationsHelper(deck, handSize, i+1, current, result)
		current = current[:len(current)-1]
	}
}
