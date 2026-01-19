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
	me := 11
	player := NewPlayerStrategyRandom(me)

	thisBoard := [BoardHeight][BoardWidth]int{
		{me, 22, me, 22, me, 22, me},
		{me, 22, me, 22, me, 22, me},
		{22, me, 22, me, 22, me, 22},
		{22, me, 22, me, 22, me, 22},
		{me, 22, me, 22, me, 22, me},
		{me, 22, me, 22, me, 22, me},
	}
	gameBoard := NewInProgressGameBoard(thisBoard)

	chosenColumn := player.PlayerChoosesAMove(gameBoard)

	if chosenColumn < 0 || chosenColumn >= BoardWidth {
		t.Errorf(`TestRandomPlayerChoosesAMoveOnEmptyBoard expected a column value within 0 - %v but played in %v column`, BoardWidth, chosenColumn)
	}
}
