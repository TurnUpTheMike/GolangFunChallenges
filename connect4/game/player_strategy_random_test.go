package game

import (
	"testing"
)

func TestRandomPlayerChoosesAMoveOnEmptyBoard(t *testing.T) {
	player := NewPlayerStrategyRandom(1)
	gameBoard := NewGameBoard()

	chosenColumn := player.PlayerChoosesAMove(*gameBoard)

	if chosenColumn < 0 || chosenColumn >= BoardWidth {
		t.Errorf(`TestRandomPlayerChoosesAMoveOnEmptyBoard expected a column value within 0 - %v but played in %v column`, BoardWidth, chosenColumn)
	}
}

func TestRandomPlayerChoosesAMoveOnFullBoard(t *testing.T) {
	player := NewPlayerStrategyRandom(1)

	thisBoard := [BoardHeight][BoardWidth]int{
		{1, 2, 1, 2, 1, 2, 1},
		{1, 2, 1, 2, 1, 2, 1},
		{2, 1, 2, 1, 2, 1, 2},
		{2, 1, 2, 1, 2, 1, 2},
		{1, 2, 1, 2, 1, 2, 1},
		{1, 2, 1, 2, 1, 2, 1},
	}
	gameBoard := NewGameBoardState(thisBoard)

	chosenColumn := player.PlayerChoosesAMove(gameBoard)

	if chosenColumn < 0 || chosenColumn >= BoardWidth {
		t.Errorf(`TestRandomPlayerChoosesAMoveOnEmptyBoard expected a column value within 0 - %v but played in %v column`, BoardWidth, chosenColumn)
	}
}
