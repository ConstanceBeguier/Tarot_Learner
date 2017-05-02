package tarot

import (
	"math/rand"
	"time"
)

type Card struct {
	Color  Color `json:"color"`
	Number int   `json:"number"`
}

type Color int

const (
	HEART Color = iota
	CLUB
	DIAMOND
	SPADE
	TRUMP
	EXCUSE
)

var nbCardsPerColor = map[Color]int{
	HEART:   14,
	CLUB:    14,
	DIAMOND: 14,
	SPADE:   14,
	TRUMP:   21,
	EXCUSE:  1,
}

func allCards() []Card {
	var cards []Card
	for col, nb := range nbCardsPerColor {
		for j := 1; j <= nb; j++ {
			cards = append(cards, Card{Color: col, Number: j})
		}
	}
	return cards
}

func random(cards []Card) []Card {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	return cards
}

func compareCards(c1 Card, c2 Card, col Color) bool {
	switch c1.Color {
	case TRUMP:
		if c2.Color == TRUMP {
			return c1.Number > c2.Number
		} else {
			return true
		}
	case col:
		if c2.Color == TRUMP {
			return false
		}
		if c2.Color == col {
			return c1.Number > c2.Number
		}
		return true
	default:
		return false
	}
}

func (c *Card) point() float32 {
	switch c.Color {
	case TRUMP:
		if c.Number == 1 || c.Number == 21 {
			return 4.5
		} else {
			return 0.5
		}
	case EXCUSE:
		return 0
	default:
		if c.Number < 11 {
			return 0.5
		} else {
			return float32(c.Number) - 9.5
		}
	}
}
