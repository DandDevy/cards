package models

import "fmt"

var kindTypes = [...]string {"spades", "hearts", "diamond"}

const (
	minNumber int8 = 1
	maxNumber = 9
)
type Card struct {
	kind string
	number int8
}

func NewCard(kind string, number int8) *Card {
	if minNumber <= number && number <= maxNumber{

		for _, element := range kindTypes{
			if kind == element{
				return &Card{kind: kind, number: number}
				break
			}
		}

	}
	return nil
}

func (c *Card) String() string {
	return fmt.Sprintf("Card{kind=%s, number=%v}",c.kind,c.number)
}

