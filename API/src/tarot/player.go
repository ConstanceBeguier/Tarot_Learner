package tarot

import (
	"fmt"
	"sort"
)

type Player struct {
	Id             int
	CardsRemaining map[Card]bool
}

type PlayerJson struct {
	AllCards []Card `json:"cards"`
	Heart    []int  `json:"0"`
	Club     []int  `json:"1"`
	Diamond  []int  `json:"2"`
	Spade    []int  `json:"3"`
	Trump    []int  `json:"4"`
	Excuse   []int  `json:"5"`
}

func (p *Player) hasCard(c Card) bool {
	notAlreadyPlayed, hasCard := p.CardsRemaining[c]
	if notAlreadyPlayed && hasCard {
		return true
	}
	return false
}

func (p *Player) validCard(c Card, t Table) bool {
	// First player can play any card
	if p.Id == t.FirstPlayer || c.Color == EXCUSE {
		return true
	}
	// Find color of this trick
	colorTrick := t.Cards[t.FirstPlayer].Color
	if colorTrick == EXCUSE {
		if p.Id == (t.FirstPlayer+1)%NB_PLAYERS {
			return true
		}
		colorTrick = t.Cards[(t.FirstPlayer+1)%NB_PLAYERS].Color
	}
	// Check the card played is valid
	switch colorTrick {
	case TRUMP:
		if !p.hasTrumps() {
			return true
		}
		if c.Color != TRUMP {
			return false
		}
		// has trumps and play a trump card
		if !p.hasBiggerTrump(t) {
			return true
		}
		for i := 0; i < NB_PLAYERS; i++ {
			if t.Cards[i].Color == TRUMP && c.Number < t.Cards[i].Number {
				return false
			}
		}
		return true
	default:
		if p.hasColor(colorTrick) {
			if c.Color == colorTrick {
				return true
			} else {
				return false
			}
		} else {
			if p.hasTrumps() {
				if c.Color != TRUMP {
					return false
				}
				if !p.hasBiggerTrump(t) {
					return true
				}
				for i := 0; i < NB_PLAYERS; i++ {
					if t.Cards[i].Color == TRUMP && c.Number < t.Cards[i].Number {
						return false
					}
				}
			}
			return true
		}
	}
}

func (p *Player) hasBiggerTrump(t Table) bool {
	maxTrump := 0
	for _, c := range t.Cards {
		if c.Color == TRUMP && maxTrump < c.Number {
			maxTrump = c.Number
		}
	}
	for c, b := range p.CardsRemaining {
		if b && c.Color == TRUMP && maxTrump < c.Number {
			return true
		}
	}
	return false
}

func (p *Player) hasTrumps() bool {
	for c, b := range p.CardsRemaining {
		if b && c.Color == TRUMP {
			return true
		}
	}
	return false
}

func (p *Player) hasColor(color Color) bool {
	for c, b := range p.CardsRemaining {
		if b && c.Color == color {
			return true
		}
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
	pl.AllCards = make([]Card, 0)
	pl.Heart = make([]int, 0)
	pl.Club = make([]int, 0)
	pl.Diamond = make([]int, 0)
	pl.Spade = make([]int, 0)
	pl.Trump = make([]int, 0)
	pl.Excuse = make([]int, 0)
	for c, b := range p.CardsRemaining {
		if b {
			pl.AllCards = append(pl.AllCards, c)
			switch c.Color {
			case HEART:
				pl.Heart = append(pl.Heart, c.Number)
			case CLUB:
				pl.Club = append(pl.Club, c.Number)
			case DIAMOND:
				pl.Diamond = append(pl.Diamond, c.Number)
			case SPADE:
				pl.Spade = append(pl.Spade, c.Number)
			case TRUMP:
				pl.Trump = append(pl.Trump, c.Number)
			case EXCUSE:
				pl.Excuse = append(pl.Excuse, c.Number)
			}
		}
	}
	sort.Ints(pl.Heart)
	sort.Ints(pl.Club)
	sort.Ints(pl.Diamond)
	sort.Ints(pl.Spade)
	sort.Ints(pl.Trump)
	sort.Ints(pl.Excuse)
	return pl
}
