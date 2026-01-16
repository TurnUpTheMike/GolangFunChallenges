package game

import (
	"errors"
	"fmt"
)

type GameBoardActions interface {
	AvailableRow(column int) int
	GetSpaceOwnership(column int, row int) int
	GetTurnHistory() []RecordedTurn
	IsPlayersSpace(player PlayerStrategy, column int, row int) bool
	IsVictory() int
	PrintGameBoard(turn int)
}

type GameBoard struct {
	board       [BoardWidth][BoardHeight]int
	turnHistory []RecordedTurn
}

type RecordedTurn struct {
	PlayerValue int
	Column      int
	Row         int
}

func NewGameBoard() *GameBoard {
	// [x][y] board coordinates, the value is player ownership
	gameBoard := [BoardWidth][BoardHeight]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
	}

	return &GameBoard{board: gameBoard, turnHistory: []RecordedTurn{}}
}

func NewGameBoardState(matrix [BoardHeight][BoardWidth]int) *GameBoard {
	// [x][y] board coordinates, the value is player ownership
	//
	// For testing readability, it is easier to visually read a transposed matrix
	return &GameBoard{board: TransposeMatrix(matrix), turnHistory: []RecordedTurn{}}
}

func NewDynamicGameState(matrix [][]int) (*GameBoard, error) {
	// [x][y] board coordinates, the value is player ownership
	//
	// This constructor allows the caller to set the gameboard using either
	// a [BoardWidth][BoardHeight] or [BoardHeight][BoardWidth] array
	if len(matrix) == BoardWidth && len(matrix[0]) == BoardHeight {
		var fixedSizeMatrix [BoardWidth][BoardHeight]int
		for i := 0; i < len(matrix); i++ {
			copy(fixedSizeMatrix[i][:], matrix[i])
		}

		return &GameBoard{board: fixedSizeMatrix, turnHistory: []RecordedTurn{}}, nil
	}

	if len(matrix) == BoardHeight && len(matrix[0]) == BoardWidth {
		var fixedSizeMatrix [BoardHeight][BoardWidth]int
		for i := 0; i < len(matrix); i++ {
			copy(fixedSizeMatrix[i][:], matrix[i])
		}

		return &GameBoard{board: TransposeMatrix(fixedSizeMatrix), turnHistory: []RecordedTurn{}}, nil
	}

	return nil, errors.New("The input must either be [BoardHeight][BoardWidth]int or [BoardWidth][BoardHeight]int")
}

func TransposeMatrix(matrix [BoardHeight][BoardWidth]int) [BoardWidth][BoardHeight]int {
	// This function is used to help visualize the board in code
	// Pass the resulting Transposed Matrix to the constructor of the GameBoard
	rows := len(matrix)
	cols := len(matrix[0])

	var transposed [BoardWidth][BoardHeight]int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

func (gameboard GameBoard) AvailableRow(column int) int {
	// Returns a StatusRowIsFull if the xSlot is full
	height := BoardHeight

	if gameboard.board[column][0] != NoPlayer {
		return StatusRowIsFull // The column is full
	}

	for y := range height {
		if y+1 >= height {
			return y
		}

		if gameboard.board[column][y+1] != NoPlayer {
			return y
		}
	}

	return StatusRowIsFull
}

func (gameBoard GameBoard) GetSpaceOwnership(column int, row int) int {
	// returns NoPlayer if owned by neither player
	return gameBoard.board[column][row]
}

func (gameBoard GameBoard) GetTurnHistory() []RecordedTurn {
	return gameBoard.turnHistory
}

func (gameBoard GameBoard) IsPlayersSpace(player PlayerStrategy, column int, row int) bool {
	return player.GetPlayerValue() == gameBoard.board[column][row]
}

func (gameBoard GameBoard) IsVictory() int {
	winner := gameBoard.IsHorizontalVictory()
	if winner != NoPlayer {
		return winner
	}

	winner = gameBoard.IsVerticalVictory()
	if winner != NoPlayer {
		return winner
	}

	winner = gameBoard.IsDiagonalVictory()
	if winner != NoPlayer {
		return winner
	}

	return NoPlayer
}

func (gameBoard GameBoard) IsHorizontalVictory() int {
	for row := range BoardHeight {
		owner := gameBoard.IsHorizontalVictoryInRow(row)

		if owner != NoPlayer {
			return owner
		}
	}

	return NoPlayer
}

func (gameBoard GameBoard) IsHorizontalVictoryInRow(row int) int {
	piecesInARow := 0
	prevSpaceOwner := NoPlayer

	for x := range BoardWidth {
		owner := gameBoard.board[x][row]

		if owner == NoPlayer {
			prevSpaceOwner = NoPlayer
			piecesInARow = 0
			continue
		}

		if owner == prevSpaceOwner {
			piecesInARow++
			if piecesInARow >= 4 {
				return owner
			}
		} else {
			piecesInARow = 1
			prevSpaceOwner = owner
		}
	}

	return NoPlayer
}

func (gameBoard GameBoard) IsVerticalVictory() int {
	for column := range BoardWidth {
		owner := gameBoard.IsVerticalVictoryInColumn(column)

		if owner != NoPlayer {
			return owner
		}
	}

	return NoPlayer
}

func (gameBoard GameBoard) IsVerticalVictoryInColumn(column int) int {
	piecesInARow := 0
	prevSpaceOwner := NoPlayer

	for y := range BoardHeight {
		owner := gameBoard.board[column][y]

		if owner == NoPlayer {
			prevSpaceOwner = NoPlayer
			piecesInARow = 0
			continue
		}

		if owner == prevSpaceOwner {
			piecesInARow++
			if piecesInARow >= 4 {
				return owner
			}
		} else {
			piecesInARow = 1
			prevSpaceOwner = owner
		}
	}

	return NoPlayer
}

func (gameBoard GameBoard) IsDiagonalVictory() int {

	for x := range BoardWidth - WinningLength - 1 {
		winner := gameBoard.IsDiagonalVictoryDownRightLane(x, 0)

		if winner != NoPlayer {
			return winner
		}
	}

	for y := 1; y < BoardWidth-WinningLength-1; y++ {
		winner := gameBoard.IsDiagonalVictoryDownRightLane(0, y)

		if winner != NoPlayer {
			return winner
		}
	}

	for x := WinningLength - 1; x < BoardWidth; x++ {
		winner := gameBoard.IsDiagonalVictoryDownLeftLane(x, 0)

		if winner != NoPlayer {
			return winner
		}
	}

	for y := 1; y < BoardHeight-(WinningLength-1); y++ {
		winner := gameBoard.IsDiagonalVictoryDownLeftLane(BoardWidth-1, y)

		if winner != NoPlayer {
			return winner
		}
	}

	return NoPlayer
}

func (gameBoard GameBoard) IsDiagonalVictoryDownLeftLane(xStart int, yStart int) int {
	spacesToCheck := min(xStart+1, BoardHeight-yStart)
	piecesInARow := 0
	prevSpaceOwner := NoPlayer

	for spaceNdx := range spacesToCheck {
		owner := gameBoard.board[xStart-spaceNdx][yStart+spaceNdx]

		if owner == NoPlayer {
			prevSpaceOwner = NoPlayer
			piecesInARow = 0
			continue
		}

		if owner == prevSpaceOwner {
			piecesInARow++
			if piecesInARow >= WinningLength {
				return owner
			}
		} else {
			piecesInARow = 1
			prevSpaceOwner = owner
		}
	}

	return NoPlayer
}

func (gameBoard GameBoard) IsDiagonalVictoryDownRightLane(xStart int, yStart int) int {
	spacesToCheck := min(BoardWidth-xStart, BoardHeight-yStart)
	piecesInARow := 0
	prevSpaceOwner := NoPlayer

	for spaceNdx := range spacesToCheck {
		owner := gameBoard.board[xStart+spaceNdx][yStart+spaceNdx]

		if owner == NoPlayer {
			prevSpaceOwner = NoPlayer
			piecesInARow = 0
			continue
		}

		if owner == prevSpaceOwner {
			piecesInARow++
			if piecesInARow >= WinningLength {
				return owner
			}
		} else {
			piecesInARow = 1
			prevSpaceOwner = owner
		}
	}

	return NoPlayer
}

func (gameBoard *GameBoard) PlayPiece(playerValue int, column int) error {
	// If an invalid row is requested to be played, a FirstAvailableMove is used for the move

	row := StatusRowIsFull
	if column > 0 && column < BoardWidth {
		row = gameBoard.AvailableRow(column)
	} // else the requested slot is out of bounds

	if row == StatusRowIsFull {
		defaultStrategy := NewPlayerStrategyFirstAvailableMove(playerValue)
		column = defaultStrategy.PlayerChoosesAMove(*gameBoard)
		if column == StatusNoAvailableMove {
			return errors.New("no available move")
		}
		row = gameBoard.AvailableRow(column)
	}

	gameBoard.board[column][row] = playerValue

	thisTurn := RecordedTurn{PlayerValue: playerValue, Column: column, Row: row}
	gameBoard.turnHistory = append(gameBoard.turnHistory, thisTurn)

	return nil
}

func (gameBoard GameBoard) PrintGameBoard(turn int) {
	width := BoardWidth
	height := BoardHeight

	for y := range height {
		fmt.Print("|  ")
		for x := range width {
			owner := gameBoard.board[x][y]

			if owner == NoPlayer {
				fmt.Print(`_  `)
			} else {
				fmt.Printf(`%d  `, owner)
			}
		}
		fmt.Println("|")
	}
	fmt.Println("|-----------------------|")

	spacing := ""
	if turn < 10 {
		spacing = " "
	}
	fmt.Printf("|        Turn  %v%d       |", spacing, turn)
	fmt.Println("")
	fmt.Println("")
}
