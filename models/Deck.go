/*
 * Created by Daniel Ashcroft on the  12 1 2020.
 */

package models

import "fmt"

const maxDeckSize int = 52
// A Deck of Cards
type Deck struct {
	cards []*Card
}



// Creates an empty instance of a Deck
func NewEmptyDeck() *Deck {
	deck := &Deck{}
	deck.cards = make([]*Card,0)
	return deck
}

// Creates an instance of a Deck from the cards given
func NewDeck(cards []*Card) *Deck {
	return &Deck{cards:cards}
}

// Creates an instance of a Deck
func NewFullDeck() *Deck {
	deck := &Deck{}
	deck.cards = make([]*Card,52) // A fixed array be better

	 for _, suit := range SuitTypes{
	 	for _, value := range ValueTypes{
			fmt.Println(value, "of", suit)
		}

	 }

	return deck
}

// Adds a Card to the Deck
func (deck *Deck) AddCard(card *Card) {
	if len(deck.cards) < maxDeckSize {
		deck.cards = append(deck.cards, card)
	}
}

// Gets all the Cards from the deck
func (deck *Deck) Cards() []*Card {
	return deck.cards
}

// Sets all the cards in the deck if size is 52 for the new one
func (deck *Deck) SetCards(cards []*Card) {
	if len(cards) == maxDeckSize {
		deck.cards = cards
	}
}

func (deck *Deck) String() string {
	return fmt.Sprintf("Deck{%s}", deck.cards)
}