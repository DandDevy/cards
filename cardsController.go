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
	deck := models.NewEmptyDeck()
	for i := 1; i < 500; i++{
		deck.AddCard(card)
	}
	deck.AddCard(card)
	fmt.Println(deck)
	fmt.Println(len(deck.Cards()))
	fmt.Println(models.SuitTypes)

	deck = models.NewFullDeck()
}
