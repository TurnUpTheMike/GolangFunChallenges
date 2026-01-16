package game

import (
	"testing"
)

func TestBlockerEmptyBoard(t *testing.T) {
	player := NewPlayerStrategyBlocker(1)
	gameBoard := NewGameBoard()

	chosenColumn := player.PlayerChoosesAMove(*gameBoard)

	// with nothing else on the board, we should always get the center column
	if chosenColumn != 3 {
		t.Errorf(`TestBlockerEmptyBoard expected a column value of 3 but got %v column`, chosenColumn)
	}
}

func TestBlockerThreeInARow(t *testing.T) {
	me := 11
	player := NewPlayerStrategyBlocker(me)
	expected := 4

	// sooooo ... this isn't really a run.
	// we'd need to keep track if things are all in a row
	textBoard := [BoardHeight][BoardWidth]int{
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, 99, 99, -1, -1, -1},
		{-1, 99, me, me, -1, -1, -1},
	}

	board := NewGameBoardState(textBoard)
	chosenColumn := player.PlayerChoosesAMove(board)

	// we should be seeing the block here!
	if chosenColumn != expected {
		t.Errorf(`TestBlockerThreeInARow expected a column value of %v but got %v column`, expected, chosenColumn)
	}
}
