package components

import "fmt"

func GetNewGame(players []*Player, size int) *Game {
	return &Game{
		players: players,
		board:   GetNewBoard(size),
	}
}

type Game struct {
	players []*Player
	board   *Board
}

func (g *Game) StartGame() {
	numPlayers, turn := len(g.players), 0
	g.board.DisplayBoard()

	for g.board.AreCellsAvailable() {
		player := g.players[turn]

		fmt.Println("-----", player.GetName(), "turn -----")

		var row int
		fmt.Println("Enter row: ")
		fmt.Scanln(&row)

		var col int
		fmt.Println("Enter column: ")
		fmt.Scanln(&col)

		if !g.board.MakeMove(row, col, player.GetPlayingPiece()) {
			fmt.Println("invalid move!!! Try again.")
			continue
		}

		g.board.DisplayBoard()

		if g.board.HasWon() {
			fmt.Println(player.GetName(), "has won")
			return
		}

		turn = (turn + 1) % numPlayers
	}

	fmt.Println("Game tied")
}
