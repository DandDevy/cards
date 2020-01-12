/*
 * Created by Daniel Ashcroft on the  12 1 2020.
 */

package models

import "fmt"

// A Deck of Cards
type Deck struct {
	cards []*Card
}



// Creates an instance of a Deck
func NewDeck() *Deck {
	return &Deck{}
}

// Adds a Card to the Deck
func (deck *Deck) AddCard(card *Card) {
	deck.SetCards(append(deck.cards, card))
}

// Gets all the Cards from the deck
func (deck *Deck) Cards() []*Card {
	return deck.cards
}

// Sets all the cards in the deck
func (deck *Deck) SetCards(cards []*Card) {
	deck.cards = cards
}

func (deck *Deck) String() string {
	return fmt.Sprintf("Deck{%s}", deck.cards)
}