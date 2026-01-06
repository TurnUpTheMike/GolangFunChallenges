package game

import (
	"testing"
)

func TestPlayConnect4(t *testing.T) {
	config := NewDefaultGameConfig()
	config.Player1 = "firstavailable"
	config.Player2 = "firstavailable"
	winner, message := PlayConnect4(config)

	if winner != 1 {
		t.Errorf(`message: %v`, message)
	}
}

func TestGetPlayerStrategyByValueOfAPlayer(t *testing.T) {
	playerFirstAvailable := NewPlayerStrategyFirstAvailableMove(4)
	playerRandom := NewPlayerStrategyRandom(5)

	players := [NumPlayers]PlayerStrategy{playerFirstAvailable, playerRandom}

	foundPlayer := GetPlayerStrategyByValue(players, 4)

	if foundPlayer == nil {
		t.Errorf(`TestGetPlayerStrategyByValueOfAPlayer did not find any player, but expected to find %d`, playerFirstAvailable.GetPlayerValue())
	}

	if foundPlayer.GetPlayerValue() != playerFirstAvailable.GetPlayerValue() {
		t.Errorf(`TestGetPlayerStrategyByValueOfAPlayer expected to find playerValue %d but found %d instead`, playerFirstAvailable.GetPlayerValue(), foundPlayer.GetPlayerValue())
	}
}

func TestGetPlayerStrategyByValueNotFound(t *testing.T) {
	playerFirstAvailable := NewPlayerStrategyFirstAvailableMove(4)
	playerRandom := NewPlayerStrategyRandom(5)

	players := [NumPlayers]PlayerStrategy{playerFirstAvailable, playerRandom}

	foundPlayer := GetPlayerStrategyByValue(players, 9)

	if foundPlayer != nil {
		t.Errorf(`TestGetPlayerStrategyByValueOfAPlayer expected to find playerValue %d but found %d instead`, playerFirstAvailable.GetPlayerValue(), foundPlayer.GetPlayerValue())
	}
}

func TestGetPlayerStrategyWithOptionString(t *testing.T) {
	playerStrategy := GetPlayerStrategy("firstavailable", 2)

	if playerStrategy.GetName() != "First Available Move Strategy" {
		t.Errorf(`TestGetPlayerStrategyWithOptionString expected to create a "First Available Move Strategy", but created %v instead`, playerStrategy.GetName())
	}

	if playerStrategy.GetPlayerValue() != 2 {
		t.Errorf(`TestGetPlayerStrategyWithOptionString expected to create a playerValue of 2, but created %d instead`, playerStrategy.GetPlayerValue())
	}
}

func TestGetPlayerStrategyWithNotFoundOption(t *testing.T) {
	playerStrategy := GetPlayerStrategy("notfound", 2)

	if playerStrategy.GetName() != "Random Move Strategy" {
		t.Errorf(`TestGetPlayerStrategyWithNotFoundOption expected to default to "Random Move Strategy", but created %v instead`, playerStrategy.GetName())
	}

	if playerStrategy.GetPlayerValue() != 2 {
		t.Errorf(`TestGetPlayerStrategyWithNotFoundOption expected to create a playerValue of 2, but created %d instead`, playerStrategy.GetPlayerValue())
	}
}
