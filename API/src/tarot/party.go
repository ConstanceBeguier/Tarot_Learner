package tarot

import ()

const NB_PLAYERS = 3
const NB_CARDS_PER_PLAYER = 24
const NB_CARDS_IN8_DOG = 6

type Party struct {
	Players [NB_PLAYERS]Player
	Dog     []Card
	Table   Table
}

func NewParty() Party {
	var party Party
	allCards := allCards()
	allCards = random(allCards)
	for i := 0; i < NB_PLAYERS; i++ {
		cards := make(map[Card]bool)
		for _, c := range allCards[NB_CARDS_PER_PLAYER*i : NB_CARDS_PER_PLAYER*(i+1)-1] {
			cards[c] = true
		}
		player := Player{
			CardsRemaining: cards,
		}
		party.Players[i] = player
	}
	party.Dog = allCards[NB_CARDS_PER_PLAYER*NB_PLAYERS:]
	party.Table = Table{}
	return party
}

func (p *Party) PlayCard(c Card, i int) bool {
	if i > NB_PLAYERS-1 || i < 0 {
		panic("Bad player number")
	}
	// check turn and player
	if !p.Table.checkTurn(i) {
		return false
	}
	if !p.Players[i].hasCard(c) {
		return false
	}
	p.Table.playCard(c, i)
	p.Players[i].removeCard(c)
	return true
}
