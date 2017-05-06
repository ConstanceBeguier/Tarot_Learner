package tarot

type Table struct {
	Scores      [2]float32       `json:"scores"`
	Cards       [NB_PLAYERS]Card `json:"cards"`
	PlayerTurn  int              `json:"playerTurn"`
	FirstPlayer int              `json:"firstPlayer"`
	TrickNb     int              `json:"trickNb"`
	IsTaker     [NB_PLAYERS]int  `json:"isTaker"`
}

func (t *Table) checkPlayerTurn(i int) bool {
	return t.PlayerTurn == i
}

func (t *Table) playCard(c Card, i int) {
	t.Cards[i] = c
	t.PlayerTurn = (t.PlayerTurn + 1) % NB_PLAYERS
	// Check end of turn
	if t.PlayerTurn != t.FirstPlayer {
		return
	}
	t.endRound()
}

func (t *Table) endRound() {
	var trickWinner int
	var trickScore float32
	// Select trick winner
	trickWinner = t.selectTrickWinner()
	// Update scores
	trickScore = t.trickScore()
	t.Scores[t.IsTaker[trickWinner]] += trickScore
	excusePlayed, playerExcuse := t.excusePlayer()
	if excusePlayed {
		t.Scores[t.IsTaker[playerExcuse]] += 4.5
	}
	// Remove cards on table
	nilCard := Card{Color: 0, Number: 0}
	for i := range t.Cards {
		t.Cards[i] = nilCard
	}
	// Select new first player
	t.FirstPlayer = trickWinner
	t.PlayerTurn = trickWinner
	t.TrickNb++
}

func (t *Table) selectTrickWinner() int {
	trickColor := t.Cards[t.FirstPlayer].Color
	trickWinner := 0
	trickBestCard := t.Cards[0]
	for i := 1; i < NB_PLAYERS; i++ {
		if compareCards(t.Cards[i], trickBestCard, trickColor) {
			trickWinner = i
			trickBestCard = t.Cards[i]
		}
	}
	return trickWinner
}

func (t *Table) trickScore() float32 {
	var score float32 = 0
	for _, c := range t.Cards {
		score += c.point()
	}
	return score
}

func (t *Table) excusePlayer() (excusePlayed bool, excusePlayer int) {
	excusePlayed = false
	excusePlayer = 0
	for i, c := range t.Cards {
		if c.Color == EXCUSE {
			excusePlayed = true
			excusePlayer = i
		}
	}
	return excusePlayed, excusePlayer
}
