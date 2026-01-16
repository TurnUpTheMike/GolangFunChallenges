package game

import (
	"errors"
	"reflect"
	"testing"
)

func TestAvailableRowOnEmptyBoard(t *testing.T) {
	gameBoard := NewGameBoard()

	availableRow := gameBoard.AvailableRow(0)

	if availableRow != BoardHeight-1 {
		t.Errorf(`TestAvailableRowOnEmptyBoard expected to return bottom row but returned %v instead`, availableRow)
	}
}

func TestAvailableRowOnPartiallyFullRow(t *testing.T) {
	thisBoard := [BoardHeight][BoardWidth]int{
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{1, -1, -1, -1, -1, -1, -1},
		{1, -1, -1, -1, -1, -1, -1},
		{2, -1, -1, -1, -1, -1, -1},
		{2, -1, -1, -1, -1, -1, -1},
	}

	gameBoard := NewGameBoardState(thisBoard)

	availableRow := gameBoard.AvailableRow(0)

	if availableRow != 1 {
		t.Errorf(`TestAvailableRowOnEmptyBoard expected to return bottom row but returned %v instead`, availableRow)
	}
}

func TestAvailableRowOnFullRow(t *testing.T) {
	thisBoard := [BoardHeight][BoardWidth]int{
		{1, -1, -1, -1, -1, -1, -1},
		{1, -1, -1, -1, -1, -1, -1},
		{0, -1, -1, -1, -1, -1, -1},
		{0, -1, -1, -1, -1, -1, -1},
		{1, -1, -1, -1, -1, -1, -1},
		{1, -1, -1, -1, -1, -1, -1},
	}

	gameBoard := NewGameBoardState(thisBoard)

	availableRow := gameBoard.AvailableRow(0)

	if availableRow != StatusRowIsFull {
		t.Errorf(`TestAvailableRowOnEmptyBoard expected to return bottom row but returned %v instead`, availableRow)
	}
}

func TestGetSpaceOwnerShipOnEmptyBoard(t *testing.T) {
	gameBoard := NewGameBoard()

	for column := range BoardWidth {
		for row := range BoardHeight {
			playerValue := gameBoard.GetSpaceOwnership(column, row)
			if playerValue != NoPlayer {
				t.Errorf(`Expected an empty board to have no ownership, column %d row %d is owned by player %d`, column, row, playerValue)
			}
		}
	}
}

func TestGetSpaceOwnershipOnPartiallyPlayedBoard(t *testing.T) {
	thisBoard := [BoardHeight][BoardWidth]int{
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{1, -1, -1, -1, -1, -1, -1},
		{1, -1, -1, -1, -1, -1, -1},
		{2, -1, -1, -1, -1, -1, -1},
		{2, -1, -1, -1, -1, -1, -1},
	}

	gameBoard := NewGameBoardState(thisBoard)

	playerValue := gameBoard.GetSpaceOwnership(0, BoardHeight-1)
	if playerValue != 2 {
		t.Errorf(`TestGetSpaceOwnershipOnPartiallyPlayedBoard expected to find ownership on space 0 %d, owned by %d`, BoardHeight-1, playerValue)
	}

	playerValue2 := gameBoard.GetSpaceOwnership(0, BoardHeight-2)
	if playerValue2 != 2 {
		t.Errorf(`TestGetSpaceOwnershipOnPartiallyPlayedBoard expected to find ownership on space 0 %d, owned by %d`, BoardHeight-2, playerValue2)
	}

	playerValue3 := gameBoard.GetSpaceOwnership(0, BoardHeight-3)
	if playerValue3 != 1 {
		t.Errorf(`TestGetSpaceOwnershipOnPartiallyPlayedBoard expected to find ownership on space 0 %d, owned by %d`, BoardHeight-3, playerValue)
	}

	playerValue4 := gameBoard.GetSpaceOwnership(0, BoardHeight-4)
	if playerValue4 != 1 {
		t.Errorf(`TestGetSpaceOwnershipOnPartiallyPlayedBoard expected to find ownership on space 0 %d, owned by %d`, BoardHeight-4, playerValue2)
	}
}

func TestGetTurnHistoryOnNewBoard(t *testing.T) {
	gameBoard := NewGameBoard()
	turnHistory := gameBoard.GetTurnHistory()

	if len(turnHistory) > 0 {
		t.Errorf("TestGetTurnHistoryOnNewBoard expected turnHistory to be an empty array")
	}
}

func TestGetTurnHistoryOnPlayedBoard(t *testing.T) {
	gameBoard := NewGameBoard()
	playerValue := 1
	columnPlayed := 3
	gameBoard.PlayPiece(playerValue, columnPlayed)

	turnHistory := gameBoard.GetTurnHistory()
	if len(turnHistory) != 1 {
		t.Errorf("TestGetTurnHistoryOnPlayedBoard expected PlayPiece() to record a turn")
	}

	firstTurn := turnHistory[0]
	if firstTurn.PlayerValue != playerValue {
		t.Errorf(`TestGetTurnHistoryOnPlayedBoard expected firstTurn.PlayerValue %d, but got %d instead`, playerValue, firstTurn.PlayerValue)
	}
	if firstTurn.Column != columnPlayed {
		t.Errorf(`TestGetTurnHistoryOnPlayedBoard expected the playedColumn to be %d, but got %d instead`, columnPlayed, firstTurn.Column)
	}
	if firstTurn.Row != BoardHeight-1 {
		t.Errorf(`TestGetTurnHistoryOnPlayedBoard expected Row to be %d, but got %d instead`, BoardHeight-1, firstTurn.Row)
	}
}

func TestIsPlayersSpaceOnEmptyBoard(t *testing.T) {
	gameBoard := NewGameBoard()

	player := NewPlayerStrategyFirstAvailableMove(1)

	spaceOwnership := gameBoard.IsPlayersSpace(player, BoardWidth-1, BoardHeight-1)

	if spaceOwnership != false {
		t.Errorf(`TestIsPlayersSpaceOnEmptyBoard expected ownership of true`)
	}
}

func TestIsPlayersSpaceOnPartiallyPlayedBoard(t *testing.T) {
	thisBoard := [BoardHeight][BoardWidth]int{
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{1, -1, -1, -1, -1, -1, -1},
		{1, -1, -1, -1, -1, -1, -1},
		{2, -1, -1, -1, -1, -1, -1},
		{2, -1, -1, -1, -1, -1, -1},
	}

	gameBoard := NewGameBoardState(thisBoard)

	player := NewPlayerStrategyFirstAvailableMove(1)

	spaceOwnership := gameBoard.IsPlayersSpace(player, 0, 2)
	if spaceOwnership != true {
		t.Errorf(`TestIsPlayersSpaceOnPartiallyPlayedBoard expect space 0, 2 to be owned by this player`)
	}
}

func TestIsHorizontalVictoryBecauseMatchIsOnBottomRow(t *testing.T) {
	thisBoard := [BoardHeight][BoardWidth]int{
		{2, -1, -1, -1, -1, -1, -1},
		{2, -1, -1, -1, -1, -1, -1},
		{1, -1, -1, 2, -1, -1, -1},
		{1, 1, -1, 1, -1, -1, -1},
		{2, 2, 2, 1, 1, -1, -1},
		{2, 2, 2, 2, 2, 1, -1},
	}

	gameBoard := NewGameBoardState(thisBoard)

	winner := gameBoard.IsHorizontalVictory()

	if winner != 2 {
		t.Errorf(`TestIsHorizontalVictoryBecauseMatchIsOnBottomRow expected to pass`)
	}
}

func TestIsHorizontalVictoryBecauseMatchIsOn2ndBottomRow(t *testing.T) {
	thisBoard := [BoardHeight][BoardWidth]int{
		{2, -1, -1, -1, -1, -1, -1},
		{2, -1, -1, -1, -1, -1, -1},
		{1, -1, -1, 2, -1, -1, -1},
		{1, 1, -1, 1, -1, -1, -1},
		{2, 2, 2, 2, 1, -1, -1},
		{2, 2, 2, 1, 2, 1, -1},
	}

	gameBoard := NewGameBoardState(thisBoard)

	winner := gameBoard.IsHorizontalVictory()

	if winner != 2 {
		t.Errorf(`TestIsHorizontalVictoryBecauseMatchIsOnBottomRow expected to pass`)
	}
}

func TestIsVerticalVictoryBecauseMatchIsVertical(t *testing.T) {
	thisBoard := [BoardHeight][BoardWidth]int{
		{2, 1, -1, 1, -1, -1, -1},
		{2, 1, -1, 2, -1, -1, -1},
		{1, 1, -1, 2, -1, -1, -1},
		{1, 1, -1, 1, -1, -1, -1},
		{2, 2, -1, 2, -1, -1, -1},
		{2, 2, -1, 1, -1, -1, -1},
	}

	gameBoard := NewGameBoardState(thisBoard)

	winner := gameBoard.IsVerticalVictory()

	if winner != 1 {
		t.Errorf(`TestIsVerticalVictoryBecauseMatchIsVertical expected to pass`)
	}
}

func TestIsDiagonalVictoryBecauseMatchIsDownLeftAndNotOnBottom(t *testing.T) {
	thisBoard := [BoardHeight][BoardWidth]int{
		{-1, -1, -1, 1, -1, -1, -1},
		{-1, -1, 1, 2, -1, -1, -1},
		{-1, 1, 1, 2, -1, -1, -1},
		{1, 2, 2, 1, -1, -1, -1},
		{1, 1, 1, 2, -1, -1, -1},
		{1, 2, 2, 1, -1, -1, -1},
	}

	gameBoard := NewGameBoardState(thisBoard)

	winner := gameBoard.IsDiagonalVictory()

	if winner != 1 {
		t.Errorf(`TestIsDiagonalVictoryBecauseMatchIsDownLeftAndNotOnBottom expected to pass`)
	}
}

func TestIsDiagonalVictoryBecauseMatchIsDownLeftAndOnBottom(t *testing.T) {
	thisBoard := [BoardHeight][BoardWidth]int{
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, 1},
		{-1, -1, -1, -1, -1, 1, 2},
		{-1, -1, -1, -1, 1, 2, 1},
		{-1, -1, -1, 1, 2, 1, 2},
	}

	gameBoard := NewGameBoardState(thisBoard)

	winner := gameBoard.IsDiagonalVictory()

	if winner != 1 {
		t.Errorf(`TestIsDiagonalVictoryBecauseMatchIsDownLeftAndOnBottom expected to pass`)
	}
}

func TestIsDiagonalVictoryBecauseMatchIsDownRight(t *testing.T) {
	thisBoard := [BoardHeight][BoardWidth]int{
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, 1, -1, -1, -1, -1},
		{-1, -1, 2, 1, -1, -1, -1},
		{-1, -1, 2, 2, 1, -1, -1},
		{-1, -1, 2, 2, 2, 1, -1},
	}

	gameBoard := NewGameBoardState(thisBoard)

	winner := gameBoard.IsDiagonalVictory()

	if winner != 1 {
		t.Errorf(`TestIsDiagonalVictoryBecauseMatchIsDownRight expected to pass`)
	}
}

func TestIsDiagonalVictoryIsNoPlayer(t *testing.T) {
	thisBoard := [BoardHeight][BoardWidth]int{
		{1, -1, -1, -1, -1, -1, -1},
		{2, 1, -1, -1, -1, -1, -1},
		{1, 2, 2, -1, -1, -1, -1},
		{2, 1, 1, 2, -1, -1, -1},
		{1, 2, 2, 1, 1, -1, -1},
		{2, 1, 2, 2, 2, 1, -1},
	}

	gameBoard := NewGameBoardState(thisBoard)

	winner := gameBoard.IsDiagonalVictoryDownRightLane(0, 0)

	if winner != NoPlayer {
		t.Errorf(`TestIsDiagonalVictoryIsNoPlayer expected no winner`)
	}
}

func TestPlayPieceOnFullBoard(t *testing.T) {
	thisBoard := [BoardHeight][BoardWidth]int{
		{2, 1, 2, 1, 2, 1, 2},
		{2, 1, 2, 1, 2, 1, 2},
		{1, 2, 1, 2, 1, 2, 1},
		{1, 2, 1, 2, 1, 2, 1},
		{2, 1, 2, 1, 2, 1, 2},
		{2, 1, 2, 1, 2, 1, 2},
	}

	gameBoard := NewGameBoardState(thisBoard)

	err := gameBoard.PlayPiece(0, 3)
	if err == nil {
		t.Errorf(`TestPlayPieceOnFullBoard expected to error but did not`)
	}
	errorNoAvailableMove := errors.New("no available move")
	if errors.Is(err, errorNoAvailableMove) {
		t.Errorf(`TestPlayPieceOnFullBoard expected error message of %v but got %v`, errorNoAvailableMove, err)
	}
}

func TestPlayPieceOnEmptyBoard(t *testing.T) {
	gameBoard := NewGameBoard()

	err := gameBoard.PlayPiece(1, 3)
	if err != nil {
		t.Errorf(`TestPlayPieceOnEmptyBoard returned error %v`, err)
	}

	transposedExpectedBoard := [BoardHeight][BoardWidth]int{
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, 1, -1, -1, -1},
	}
	expectedBoard := TransposeMatrix(transposedExpectedBoard)
	if !reflect.DeepEqual(gameBoard.board, expectedBoard) {
		t.Errorf(`TestPlayPieceOnEmptyBoard expected to alter the board differently`)
	}
}

func TestPlayPieceOutOfBounds(t *testing.T) {
	gameBoard := NewGameBoard()

	err := gameBoard.PlayPiece(1, BoardWidth+1)
	if err != nil {
		t.Errorf(`TestPlayPieceOnEmptyBoard returned error %v`, err)
	}

	transposedExpectedBoard := [BoardHeight][BoardWidth]int{
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, 1, -1, -1, -1},
	}
	expectedBoard := TransposeMatrix(transposedExpectedBoard)
	if !reflect.DeepEqual(gameBoard.board, expectedBoard) {
		t.Errorf(`TestPlayPieceOnEmptyBoard expected to alter the board differently`)
	}
}
