package tarot

import (
	"testing"
)

func TestAllCards(t *testing.T) {
	cards := allCards()
	if len(cards) != 78 {
		t.Errorf("allCards should return 78 cards\n")
	}
}

func TestCompareCards(t *testing.T) {
	cards := make(map[Color][]Card)
	for col, nb := range nbCardsPerColor {
		if _, ok := cards[col]; !ok {
			cards[col] = make([]Card, 0)
		}
		cards[col] = append(cards[col], Card{Color: col, Number: nb})
	}
	// TRUMP is greater than other colors
	var col Color
	var trickColor Color
	for col = 0; col < 6; col++ {
		if col == TRUMP {
			continue
		}
		for _, c1 := range cards[col] {
			for _, c2 := range cards[TRUMP] {
				for trickColor = 0; trickColor < 5; trickColor++ {
					if compareCards(c1, c2, trickColor) {
						t.Errorf("Problem: %v should be geater than %v", c2, c1)
					}
					if !compareCards(c2, c1, trickColor) {
						t.Errorf("Problem: %v should be geater than %v", c2, c1)
					}
				}
			}
		}
	}
	// EXCUSE cannot win the trick
	excuseCard := Card{Color: EXCUSE, Number: 1}
	for col = 0; col < 6; col++ {
		if col == EXCUSE {
			continue
		}
		for _, c1 := range cards[col] {
			for trickColor = 0; trickColor < 5; trickColor++ {
				if compareCards(excuseCard, c1, trickColor) {
					t.Errorf("Problem: %v should be geater than %v", c1, excuseCard)
				}
			}
		}
	}
}
