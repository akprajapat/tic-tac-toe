package main

import (
	"github.com/akprajapat/tic-tac-toe/game"
)

func main() {
	g := game.NewGame()
	g.PrintPlayers()
	g.PlayGame()
}
