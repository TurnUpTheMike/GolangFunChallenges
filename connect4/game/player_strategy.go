package game

type PlayerStrategy interface {
	PlayerChoosesAMove(GameBoardActions) int
	GetName() string
	GetPlayerValue() int
}

type PlayerStrategyFirstAvailableMove struct {
	name        string
	playerValue int
}

func NewPlayerStrategyFirstAvailableMove(playerValue int) *PlayerStrategyFirstAvailableMove {
	return &PlayerStrategyFirstAvailableMove{
		name:        "First Available Move Strategy",
		playerValue: playerValue,
	}
}

func (p PlayerStrategyFirstAvailableMove) GetName() string {
	return p.name
}

func (p PlayerStrategyFirstAvailableMove) GetPlayerValue() int {
	return p.playerValue
}

func (p PlayerStrategyFirstAvailableMove) PlayerChoosesAMove(gameBoard GameBoardActions) int {
	columnPriority := [BoardWidth]int{3, 2, 4, 1, 5, 0, 6}

	for _, column := range columnPriority {
		if gameBoard.AvailableRow(column) != StatusRowIsFull {
			return column
		}
	}

	return StatusNoAvailableMove
}
