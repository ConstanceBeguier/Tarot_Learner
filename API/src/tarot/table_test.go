package tarot

import (
	"testing"
)

func TestCheckTurn(t *testing.T) {
	table := Table{
		Turn: 2,
	}
	if table.checkTurn(0) {
		t.Errorf("Problem with turn function\n")
	}
	if table.checkTurn(1) {
		t.Errorf("Problem with turn function\n")
	}
	if !table.checkTurn(2) {
		t.Errorf("Problem with turn function\n")
	}

}
