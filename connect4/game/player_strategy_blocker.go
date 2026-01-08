package game

type PlayerStrategyBlocker struct {
	playerValue int
}

func NewPlayerStrategyBlocker(playerValue int) *PlayerStrategyBlocker {
	return &PlayerStrategyBlocker{
		playerValue: playerValue,
	}
}

func (p PlayerStrategyBlocker) GetName() string {
	return "Blocking Strategy"
}

func (p PlayerStrategyBlocker) GetPlayerValue() int {
	return p.playerValue
}

func (p PlayerStrategyBlocker) PlayerChoosesAMove(gba GameBoardActions) int {
	// check the board for any three-in-row
	// block that move if can

	running := 0
	blocker := []int{}

	// just start looking from the first column to the next-to-last
NextColumnLoop:
	for i := 0; i < (BoardWidth - 1); i++ {
		for j := 0; j < BoardHeight; j++ {
			owner := gba.GetSpaceOwnership(i, j)
			switch owner {
			case NoPlayer:
				// this is empty, keep going
				continue
			case p.playerValue:
				// you own it. reset running and continue to the next column
				running = 0
				continue NextColumnLoop
			default:
				// they own it; can you
				running += 1
				if running == 3 {
					blocker = append(blocker, i+1)
					break NextColumnLoop
				}
				continue NextColumnLoop
			}
		}
	}

	// see if we can drop one in there
	// we'll first try the candidate, if any, then
	// fall back to a "center first" strategy
	candidates := append(blocker, []int{3, 2, 4, 1, 5, 0, 6}...)
	for _, c := range candidates {
		if gba.AvailableRow(c) != StatusRowIsFull {
			return c
		}
	}

	return StatusNoAvailableMove
}
