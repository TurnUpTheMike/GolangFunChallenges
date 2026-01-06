package game

import (
	"math/rand"
)

type PlayerStrategyRandom struct {
	playerValue int
}

func NewPlayerStrategyRandom(playerValue int) *PlayerStrategyRandom {
	return &PlayerStrategyRandom{
		playerValue: playerValue,
	}
}

func (p PlayerStrategyRandom) GetName() string {
	return "Random Move Strategy"
}

func (p PlayerStrategyRandom) GetPlayerValue() int {
	return p.playerValue
}

func (p PlayerStrategyRandom) PlayerChoosesAMove(gameBoard GameBoardActions) int {
	column := rand.Intn(BoardWidth)
	return column
}
