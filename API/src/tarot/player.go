package tarot

import (
	"fmt"
)

type Player struct {
	CardsRemaining map[Card]bool `json:"cards,omitempty"`
}

type PlayerJson struct {
	CardsRemaining []Card `json:"cards,omitempty"`
}

func (p *Player) hasCard(c Card) bool {
	notAlreadyPlayed, hasCard := p.CardsRemaining[c]
	if notAlreadyPlayed && hasCard {
		return true
	}
	return false
}

func (p *Player) removeCard(c Card) {
	notAlreadyPlayed, hasCard := p.CardsRemaining[c]
	if notAlreadyPlayed && hasCard {
		p.CardsRemaining[c] = false
		return
	}
	panic(fmt.Errorf("This player does not have card %q\n", c))
}

func (p *Player) CardsToJson() PlayerJson {
	var pl PlayerJson
	for c, b := range p.CardsRemaining {
		if b {
			pl.CardsRemaining = append(pl.CardsRemaining, c)
		}
	}
	return pl
}
