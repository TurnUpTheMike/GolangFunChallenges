# Let's play Connect 4

Here's how to play

Create a pull request to this repo that contains
- A new file called player_strategy_<strategy_name>.go
- Modify game.go func GetPlayerStrategy to call your code's constructor and return a PlayerStrategy

Your code in player_strategy_<strategy_name>.go should implement the functions of the Interfrace PlayerStrategy found in player_strategy.go.

Your code should also contain a function NewPlayerStrategy<strategy_name> that returns your implementation of PlayerStrategy.

# Example Usage

Running with defaults
`go run .`

Selecting a specific PlayerStrategy
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

This repository uses golang's standard test runner
`go test -v`
