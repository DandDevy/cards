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
var suitTypes = [numberOfSuitTypes]string {"spades", "hearts", "diamonds", "clubs"}
var valueTypes = [numberOfValueTypes]string { "Ace", "2", "3", "4", "5","6", "7", "8", "9", "10", "Jack", "Queen", "King"}

type Card struct {
	suit  string
	value string
}

func NewCard(suit string, value string) *Card {

	var returningCard *Card

	// checks if suit is possible suit type
	for _, suitType := range suitTypes {
		if suit == suitType {

			// checks if value is possible value type
			for _, valueType := range valueTypes{
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

func (c *Card) String() string {
	return fmt.Sprintf("Card{suit=%s, value=%v}",c.suit,c.value)
}

