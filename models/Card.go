package models

var kind_types = [...]string {"spades", "hearts", "diamond"}

const (
	minNumber int8 = 0
	maxNumber = 0
)
type Card struct {
	kind string
	number int8
}

func NewCard(kind string, number int8) *Card {

	return &Card{kind: kind, number: number}
}
