package game

import (
	"testing"
)

func TestPlayerChoosesAMoveOnEmptyBoard(t *testing.T) {
	player := NewPlayerStrategyFirstAvailableMove(0)
	gameBoard := NewGameBoard()
	middleColumn := BoardWidth / 2

	chosenColumn := player.PlayerChoosesAMove(*gameBoard)

	if chosenColumn != middleColumn {
		t.Errorf(`TestPlayerChoosesAMoveOnEmptyBoard expected %v column but played in %v column`, middleColumn, chosenColumn)
	}
}

func TestPlayerChoosesAMoveOnFullMiddleRow(t *testing.T) {
	player := NewPlayerStrategyFirstAvailableMove(0)

	thisBoard := [BoardHeight][BoardWidth]int{
		{-1, -1, -1, 0, -1, -1, -1},
		{-1, -1, -1, 1, -1, -1, -1},
		{-1, -1, -1, 1, -1, -1, -1},
		{-1, -1, -1, 0, -1, -1, -1},
		{-1, -1, -1, 1, -1, -1, -1},
		{-1, -1, -1, 0, -1, -1, -1},
	}

	gameBoard := GameBoard{board: TransposeMatrix(thisBoard)}
	leftOfMiddleColumn := (BoardWidth / 2) - 1

	chosenColumn := player.PlayerChoosesAMove(gameBoard)

	if chosenColumn != leftOfMiddleColumn {
		t.Errorf(`TestPlayerChoosesAMoveOnFullMiddleRow expected %v column but played in %v column`, leftOfMiddleColumn, chosenColumn)
	}
}

func TestPlayerChoosesAMoveOnFullBoard(t *testing.T) {
	player := NewPlayerStrategyFirstAvailableMove(0)

	thisBoard := [BoardHeight][BoardWidth]int{
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1},
		{0, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0},
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1},
	}
	gameBoard := GameBoard{board: TransposeMatrix(thisBoard)}

	chosenColumn := player.PlayerChoosesAMove(gameBoard)

	if chosenColumn != StatusNoAvailableMove {
		t.Errorf(`TestPlayerChoosesAMoveOnFullBoard expected NoAvailableMoveStatus but played in %v column`, chosenColumn)
	}
}
