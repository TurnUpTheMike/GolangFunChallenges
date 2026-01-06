package main

import (
	"fmt"
	"connect4/game"
	"flag"
)

func main() {
	fmt.Println("Let's Play Connect 4")

	argPlayer1 := flag.String("player1", "firstavailable", "The Player Strategy key for Player 1")
	argPlayer2 := flag.String("player2", "firstavailable", "The Player Strategy key for Player 2")
	argPrintBoardCadence := flag.Int("printboard", 5, "Print the board to the display every n turns")
	flag.Parse()

	config := game.NewDefaultGameConfig()
	config.Player1 = *argPlayer1
	config.Player2 = *argPlayer2
	config.ModuloToPrintGameBoard = *argPrintBoardCadence

	_, message := game.PlayConnect4(config)

	fmt.Println(message)
	fmt.Println("El Fin")
}