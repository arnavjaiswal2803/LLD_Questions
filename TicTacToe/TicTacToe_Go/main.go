package main

import (
	"ticTacToe/components"
	"ticTacToe/components/playingPieces"
)

func main() {
	players := make([]*components.Player, 0)
	players = append(players, components.GetNewPlayer("Foo", playingPieces.GetPlayingPieceX()))
	players = append(players, components.GetNewPlayer("Bar", playingPieces.GetPlayingPieceO()))

	game := components.GetNewGame(players, 3)

	game.StartGame()
}
