/*
 * Created by Daniel Ashcroft on the  12 1 2020.
 */

package main

import (
	"fmt"
	"github.com/DandDevy/cards/models"
)

// Creates a new Deck
func NewDeck() {
	card := models.NewCard("hearts", "3")
	deck := models.NewDeck()
	deck.AddCard(card)
	deck.AddCard(card)
	deck.AddCard(card)
	deck.AddCard(card)
	fmt.Println(deck)
}
