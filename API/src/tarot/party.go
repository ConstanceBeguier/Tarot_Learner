package tarot

import ()

const NB_PLAYERS = 3
const NB_CARDS_PER_PLAYER = 24
const NB_CARDS_IN8_DOG = 6

type Party struct {
	Players [NB_PLAYERS]Player
	Kitty   []Card
	Table   Table
	Seats   Seats
}

type Seats struct {
	AvailableSeats [NB_PLAYERS]bool `json:"availableSeats"`
}

type ValidCardsJson struct {
	IsYourTurn bool   `json:"isYourTurn"`
	ValidCards []Card `json:"validCards"`
}

func NewParty() Party {
	var party Party
	allCards := allCards()
	allCards = random(allCards)
	for i := 0; i < NB_PLAYERS; i++ {
		cards := make(map[Card]bool)
		for _, c := range allCards[NB_CARDS_PER_PLAYER*i : NB_CARDS_PER_PLAYER*(i+1)] {
			cards[c] = true
		}
		player := Player{
			Id:             i,
			CardsRemaining: cards,
		}
		party.Players[i] = player
	}
	party.Kitty = allCards[NB_CARDS_PER_PLAYER*NB_PLAYERS:]
	party.Table = Table{IsTaker: [NB_PLAYERS]int{1, 0, 0}}
	for i := range party.Seats.AvailableSeats {
		party.Seats.AvailableSeats[i] = true
	}
	return party
}

func (p *Party) PlayCard(c Card, i int) bool {
	if i > NB_PLAYERS-1 || i < 0 {
		panic("Bad player number")
	}
	// check turn and player
	if !p.Table.checkPlayerTurn(i) {
		return false
	}
	if !p.Players[i].hasCard(c) {
		return false
	}
	if !p.Players[i].validCard(c, p.Table) {
		return false
	}
	p.Table.playCard(c, i)
	p.Players[i].removeCard(c)
	return true
}

func (p *Party) ValidCards(playerNb int) ValidCardsJson {
	if playerNb != p.Table.PlayerTurn {
		return ValidCardsJson{IsYourTurn: false, ValidCards: []Card{}}
	}
	var possibleCards []Card = make([]Card, 0)
	for c, b := range p.Players[playerNb].CardsRemaining {
		if b {
			if p.Players[playerNb].validCard(c, p.Table) {
				possibleCards = append(possibleCards, c)
			}
		}
	}
	return ValidCardsJson{IsYourTurn: true, ValidCards: possibleCards}
}
