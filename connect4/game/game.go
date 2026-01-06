package game

import (
	"errors"
	"fmt"
	"strings"
)

const WinningLength int = 4
const BoardWidth int = 7
const BoardHeight int = 6
const NumPlayers int = 2
const NoPlayer int = -1
const StatusRowIsFull int = -1
const StatusNoAvailableMove int = -2

type GameConfig struct {
	Player1                string
	Player2                string
	ModuloToPrintGameBoard int
}

func NewDefaultGameConfig() GameConfig {
	return GameConfig{
		Player1:                "random",
		Player2:                "random",
		ModuloToPrintGameBoard: 5,
	}
}

func PlayConnect4(config GameConfig) (int, string) {
	errorNoAvailableMove := errors.New("no available move")

	gameBoard := NewGameBoard()
	playerValues := [NumPlayers]int{1, 2}

	player1 := GetPlayerStrategy(config.Player1, playerValues[0])
	player2 := GetPlayerStrategy(config.Player2, playerValues[1])
	players := [NumPlayers]PlayerStrategy{player1, player2}

	winner := NoPlayer
	
	turn := 0
	for turn = range BoardWidth * BoardHeight {
		whosTurn := turn % NumPlayers
		columnChosen := players[whosTurn].PlayerChoosesAMove(gameBoard)

		err := gameBoard.PlayPiece(playerValues[whosTurn], columnChosen)

		if errors.Is(err, errorNoAvailableMove) {
			winner = NoPlayer
			break
		}

		winner = gameBoard.IsVictory()
		if winner != NoPlayer {
			break
		}

		if turn%config.ModuloToPrintGameBoard == 0 {
			gameBoard.PrintGameBoard(turn)
		}
	}

	gameBoard.PrintGameBoard(turn)

	winningPlayer := GetPlayerStrategyByValue(players, winner)

	if winningPlayer == nil {
		return NoPlayer, fmt.Sprintf(`Turn %d the Winner is Neither Player. The Match has ended in a tie`, turn)
	}

	message := fmt.Sprintf(`Turn %d the Winner is Player %d %v`, turn, winner, winningPlayer.GetName())
	return winner, message
}

func GetPlayerStrategyByValue(players [NumPlayers]PlayerStrategy, playerValue int) PlayerStrategy {
	for playerndx := range NumPlayers {
		if playerValue == players[playerndx].GetPlayerValue() {
			return players[playerndx]
		}
	}

	return nil
}

func GetPlayerStrategy(option string, playerValue int) PlayerStrategy {
	switch strings.ToLower(option) {
	case "random":
		return NewPlayerStrategyRandom(playerValue)
	case "firstavailable":
		return NewPlayerStrategyFirstAvailableMove(playerValue)
	}

	// default
	return NewPlayerStrategyRandom(playerValue)
}
