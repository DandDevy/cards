package main

import (
	"fmt"
	"github.com/DandDevy/cards/models"
)

func NewDeck() {
	card := models.NewCard("hearts", 3)
	fmt.Println(card)
}
