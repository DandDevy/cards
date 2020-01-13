/*
 * Created by Daniel Ashcroft on the  12 1 2020.
 */

package models

import "fmt"

const (
	numberOfSuitTypes = 4
	numberOfValueTypes = 13
)

// the list of possible suits and values allowed
var SuitTypes = [numberOfSuitTypes]string {"spades", "hearts", "diamonds", "clubs"}
var ValueTypes = [numberOfValueTypes]string {"Ace", "2", "3", "4", "5","6", "7", "8", "9", "10", "Jack", "Queen", "King"}

// A card that is inside of a deck with a suit and a value
type Card struct {
	suit  string
	value string
}

// If the suit and the value details are present in SuitTypes and ValueTypes then this constructor will return a card
func NewCard(suit string, value string) *Card {

	var returningCard *Card

	// checks if suit is possible suit type
	for _, suitType := range SuitTypes {
		if suit == suitType {

			// checks if value is possible value type
			for _, valueType := range ValueTypes {
				if valueType == value {
					returningCard =  &Card{suit: suit, value: value}
					break
				}
			}
			break
		}
	}

	return returningCard
}

func (c *Card) Value() string {
	return c.value
}

func (c *Card) SetValue(value string) {
	c.value = value
}

func (c *Card) Suit() string {
	return c.suit
}

func (c *Card) SetSuit(suit string) {
	c.suit = suit
}

func (c *Card) String() string {
	return fmt.Sprintf("Card{suit=%s, value=%v}",c.suit,c.value)
}

