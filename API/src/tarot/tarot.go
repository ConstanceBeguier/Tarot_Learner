package tarot

import (
	"math/rand"
	"time"
)

const NB_PLAYERS = 3
const NB_CARDS_PER_PLAYER = 24
const NB_CARDS_IN8_DOG = 6

type Card struct {
	Color  Color `json:"color,omitempty"`
	Number int   `json:"number,omitempty"`
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

type Player struct {
	cardsRemaining []Card
	role           int // 1 attack and 0 defense
}

type Party struct {
	players [NB_PLAYERS]Player
	dog     []Card
	// scores[0]: defense score
	// scores[1]: attack score
	Scores TableScores `json:"tableScores,omitempty"`
}

type TableScores struct {
	// scores[0]: defense score
	// scores[1]: attack score
	Scores [2]int
}

func CardHeart1() Card {
	return Card{Color: HEART, Number: 1}
}

func random(cards []Card) []Card {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	return cards
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

func dealing() Party {
	var party Party
	allCards := allCards()
	allCards = random(allCards)
	for i := 0; i < NB_PLAYERS; i++ {
		player := Player{
			cardsRemaining: allCards[NB_CARDS_PER_PLAYER*i : NB_CARDS_PER_PLAYER*(i+1)-1],
			role:           0,
		}
		party.players[i] = player
	}
	party.dog = allCards[NB_CARDS_PER_PLAYER*NB_PLAYERS:]
	return party
}
