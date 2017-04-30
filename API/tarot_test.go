package tarot

import (
  "fmt"
  "testing"
)

func TestAllCards(t *testing.T) {
  cards := allCards()
  fmt.Println(cards)
  cards = random(cards)
  fmt.Println(cards)
  if len(cards) != 78 {
    t.Error("Wong number of cards")
  }
}

func TestDealing(t *testing.T) {
  party := dealing()
  fmt.Println("Players")
  for i := 0 ; i<3 ; i++ {
    fmt.Println(party.players[i])
  }
  fmt.Println("Dog")
  fmt.Println(party.dog)
  fmt.Println("Scores")
  fmt.Println(party.scores)
}
