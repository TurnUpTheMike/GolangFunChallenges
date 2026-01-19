package main

import (
	"connect4/game"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Let's Play Connect 4")
	playerRegistryHelp := game.GetHelpMessageOfPlayerRegistry()

	argPlayer1 := flag.String("player1", "random", playerRegistryHelp)
	argPlayer2 := flag.String("player2", "random", playerRegistryHelp)
	argPrintBoardCadence := flag.Int("printboard", 5, "Print the board to the display every n turns")
	flag.Parse()

	config := game.NewDefaultGameConfig()
	config.Player1 = *argPlayer1
	config.Player2 = *argPlayer2
	config.ModuloToPrintGameBoard = *argPrintBoardCadence

	_, message := game.PlayConnect4(config)

	fmt.Println(message)
}
