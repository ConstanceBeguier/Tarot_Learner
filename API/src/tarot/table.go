package tarot

type Table struct {
	Scores      [2]float32       `json:"tableScores,omitempty"`
	Cards       [NB_PLAYERS]Card `json:"tableCards,omitempty"`
	Turn        int              `json:"tableTurn,omitempty"`
	FirstPlayer int              `json:"firstPlayer,omitempty"`
	IsAttacker  [NB_PLAYERS]int  `json:"isAttacker,omitempty"`
}

func (t *Table) checkTurn(i int) bool {
	return t.Turn == i
}

func (t *Table) playCard(c Card, i int) {
	t.Cards[i] = c
	t.Turn = (t.Turn + 1) % NB_PLAYERS
	if t.Turn != t.FirstPlayer {
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
	t.Scores[t.IsAttacker[trickWinner]] += trickScore
	excusePlayed, playerExcuse := t.excusePlayer()
	if excusePlayed {
		t.Scores[t.IsAttacker[playerExcuse]] += 4.5
	}
	// Remove cards on table
	// Select new first player
	t.FirstPlayer = trickWinner
	t.Turn = trickWinner
}

func (t *Table) selectTrickWinner() int {
	trickColor := t.Cards[t.Turn].Color
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
