package tarot

import (
	"fmt"
	"sort"
)

type Player struct {
	CardsRemaining map[Card]bool
}

type PlayerJson struct {
	Heart   []int `json:"heart_cards"`
	Club    []int `json:"club_cards"`
	Diamond []int `json:"diamond_cards"`
	Spade   []int `json:"spade_cards"`
	Trump   []int `json:"trump_cards"`
	Excuse  []int `json:"excuse_cards"`
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
	pl.Heart = make([]int, 0)
	pl.Club = make([]int, 0)
	pl.Diamond = make([]int, 0)
	pl.Spade = make([]int, 0)
	pl.Trump = make([]int, 0)
	pl.Excuse = make([]int, 0)
	for c, b := range p.CardsRemaining {
		if b {
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
