package tarot

import (
	"fmt"
)

type Player struct {
	CardsRemaining map[Card]bool
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
