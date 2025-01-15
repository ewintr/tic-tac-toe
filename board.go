package main

import "fmt"

type Board struct {
	// the 3x3 board is stored by concatenating all rows
	//
	// 0 1 2
	// 3 4 5
	// 6 7 8
	//
	// becomes:
	//
	// 0 1 2 3 4 5 6 7 8
	//
	// each square contains either one of the player marks, or a space
	squares [9]string
}

func NewBoard() *Board {
	s := [9]string{}
	for i := range s {
		s[i] = " "
	}
	return &Board{
		squares: s,
	}
}

// Mark puts a mark in a square.
// returns false if the square was already occupied
func (b *Board) Mark(sq int, mark string) bool {
	if b.squares[sq] != " " {
		return false
	}

	b.squares[sq] = mark
	return true
}

func (b *Board) Available() []int {
	av := make([]int, 0)
	for p, m := range b.squares {
		if m == " " {
			av = append(av, p)
		}
	}

	return av
}

func (b *Board) Full() bool {
	return len(b.Available()) == 0
}

func (b *Board) Winner() (string, bool) {
	// someone has won when the board features identical marks on all these places
	winConfs := [][3]int{
		{0, 1, 2}, // horizontal rows
		{3, 4, 5},
		{6, 7, 8},

		{0, 3, 6}, // vertical rows
		{1, 4, 7},
		{2, 5, 8},

		{0, 4, 8}, // diagonals
		{2, 4, 6},
	}

	for _, w := range winConfs {
		// first determine what mark is on the first position
		m := b.squares[w[0]]
		if m == " " {
			continue
		}
		// then check whether the others are the same
		if m == b.squares[w[1]] && m == b.squares[w[2]] {
			return m, true
		}
	}

	return "", false
}

func (b *Board) Render(useEsc bool) string {
	sqs := make([]string, 0, 8)
	for _, sq := range b.squares {
		sqStr := sq
		if useEsc {
			// make the marks bold and white
			sqStr = fmt.Sprintf("\x1b[1;37m%s\x1b[0m", sqStr)
		}
		sqs = append(sqs, sqStr)
	}

	return fmt.Sprintf(` -0--- -1--- -2---
|     |     |     |
|  %s  |  %s  |  %s  |
|     |     |     |
 -3--- -4--- -5---
|     |     |     |
|  %s  |  %s  |  %s  |
|     |     |     |
 -6--- -7--- -8---
|     |     |     |
|  %s  |  %s  |  %s  |
|     |     |     |
 ----- ----- -----
`, sqs[0], sqs[1], sqs[2],
		sqs[3], sqs[4], sqs[5],
		sqs[6], sqs[7], sqs[8])
}
