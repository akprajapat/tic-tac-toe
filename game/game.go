package game

import "fmt"

type game struct {
	players []*player
	board   *board
}

func NewGame() *game {
	num_players, board_size := 0, 0
	fmt.Printf("Enter number of Players and size of Board: ")
	fmt.Scanf("%d %d", &num_players, &board_size)
	players := make([]*player, num_players)
	for i := 1; i <= num_players; i++ {
		var name string
		fmt.Printf("Enter player%d name:", i)
		fmt.Scan(&name)
		player := newPlayer(i, name)
		players[i-1] = player
	}
	return &game{
		players: players,
		board:   newBoard(board_size),
	}
}

func (g *game) PrintPlayers() {
	for _, player := range g.players {
		fmt.Println(player.name)
	}
}

func (g *game) PlayGame() {
	r, c := 0, 0
	count, countPieces := 0, 0

	for {
		playerId := count%len(g.players) + 1
		fmt.Printf("Player%d's turn enter row col: ", playerId)
		fmt.Scanf("%d %d", &r, &c)
		if r >= g.board.size || r < 0 || c >= g.board.size || c < 0 {
			fmt.Printf("Row or Col must in Given range %d %d\n", 0, g.board.size)
			continue
		}
		isFilled := g.board.fill(playerId, r, c)

		if !isFilled {
			fmt.Println("Given row col is already filled")
			continue
		}
		countPieces += 1
		g.board.printBoard()
		if g.board.isWinner(playerId, r, c) {
			fmt.Printf("Congrats! Player%d Wins The Game.\n", playerId)
			return
		}
		if countPieces == g.board.size*g.board.size {
			fmt.Println("Both Players played well, Game Draws")
			return
		}
		count++
	}
}
