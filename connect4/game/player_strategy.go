package game

import (
	"fmt"
	"strings"
	"sync"
)

type PlayerStrategy interface {
	PlayerChoosesAMove(GameBoardActions) int
	GetName() string
	GetPlayerValue() int
}

// playerRegistry is the central map storing constructors/factories.
type PlayerStrategyFactory func(ownershipValue int) PlayerStrategy

var playerRegistry = make(map[string]PlayerStrategyFactory)
var registryMutex sync.RWMutex

// Register adds a new PlayerStrategy constructor to the registry.
func Register(optionName string, constructor PlayerStrategyFactory) {
	registryMutex.Lock()
	defer registryMutex.Unlock()
	if _, exists := playerRegistry[optionName]; exists {
		// Handle error or panic if the name is already registered
		panic(fmt.Sprintf("playerstrategy %s already registered", optionName))
	}
	playerRegistry[optionName] = constructor
}

// Get retrieves a PlayerStrategy instance by name.
func GetRegisteredPlayerStrategy(name string, playerValue int) PlayerStrategy {
	registryMutex.RLock()
	defer registryMutex.RUnlock()
	constructor, exists := playerRegistry[name]
	if !exists {
		return nil // Or return an error
	}
	return constructor(playerValue)
}

func GetHelpMessageOfPlayerRegistry() string {
	var keys []string
	for k := range playerRegistry {
		keys = append(keys, k)
	}
	message := "Options: " + strings.Join(keys, ", ")
	return message
}

// An example of an implemented PlayerStrategy
func init() {
	Register("firstavailable", NewPlayerStrategyFirstAvailableMove)
}

type PlayerStrategyFirstAvailableMove struct {
	name        string
	playerValue int
}

func NewPlayerStrategyFirstAvailableMove(playerValue int) PlayerStrategy {
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
