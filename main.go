package main

import (
	"fmt"
	"os"
)

func main() {
	useEscapeCodes := true
	if len(os.Args) == 2 && os.Args[1] == "plain" {
		useEscapeCodes = false
	}

	game := NewGame(NewHuman("X"), NewHuman("O"), useEscapeCodes)
	for {
		if done := game.Turn(); done {
			os.Exit(0)
		}
	}
}

type Game struct {
	turns   int
	players [2]Player
	board   *Board
	useEsc  bool
}

func NewGame(p1, p2 Player, useEsc bool) *Game {
	return &Game{
		players: [2]Player{p1, p2},
		board:   NewBoard(),
		useEsc:  useEsc,
	}
}

// Turn returns true when the game is finished.
func (g *Game) Turn() bool {
	g.RenderBoard()

	// get next move
	curPl := g.turns % 2
	fmt.Printf("Turn %d: player %d (%s) can make a move\n", g.turns, curPl, g.players[curPl].Mark())
	if cont := g.players[curPl].MakeMove(g.board); !cont {
		fmt.Println("Maybe some other time then. Bye!")
		return true
	}

	// check result
	if _, ok := g.board.Winner(); ok {
		g.RenderBoard()
		fmt.Printf("Congratulations player %d, you win!\n", curPl)
		return true
	}
	if g.board.Full() {
		g.RenderBoard()
		fmt.Println("Stalemate! Try again...")
		return true
	}

	g.turns++

	return false
}

func (g *Game) RenderBoard() {
	if g.useEsc {
		// clear screen first
		fmt.Print("\033[H\033[2J")
	}
	fmt.Println(g.board.Render(g.useEsc))
}
