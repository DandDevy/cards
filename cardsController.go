/*
 * Created by Daniel Ashcroft on the  12 1 2020.
 */

package main

import (
	"fmt"
	"github.com/DandDevy/cards/models"
)

func NewDeck() {
	card := models.NewCard("hearts", "3")
	fmt.Println(card)
}
