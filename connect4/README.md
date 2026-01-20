# Let's play Connect 4

Here's how to play

Create a pull request to this repo that contains
- A new file called `player_strategy_<strategy_name>.go`
- Your code in `player_strategy_<strategy_name>.go` implements the functions of the Interfrace PlayerStrategy found in `player_strategy.go`.

- Your code contains a factory Function <br/> 
`NewPlayerStrategy<strategy_name>(int playerValue) PlayerStrategy` <br/>
that returns your implementation of PlayerStrategy.

- Your code registers your PlayerStrategy with an optionValue like
```
func init() {
	Register("random", NewPlayerStrategyRandom)
}
```

# Example Usage

Running with defaults <br/>
`go run .`

Selecting a specific PlayerStrategy <br/>
`go run . --player1 random --player2 firstavailable`

```
go run . --help

  -player1 string
        The Player Strategy key for Player 1 (default "firstavailable")
  -player2 string
        The Player Strategy key for Player 2 (default "firstavailable")
  -printboard int
        Print the board to the display every n turns (default 5)
```

# Running Tests

This repository uses golang's standard test runner <br/>
`go test -v`
