package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Player interface {
	Mark() string
	MakeMove(b *Board) bool
}

type Human struct {
	mark string
}

func NewHuman(m string) *Human {
	return &Human{mark: m}
}

func (h *Human) Mark() string { return h.mark }

// MakeMove marks the board
// returns false if the player wants to stop
func (h *Human) MakeMove(b *Board) bool {
	for {
		sq, stop, err := AskInput()
		if err != nil {
			fmt.Printf("That didn't work: %v\n", err)
			fmt.Println("Try again, or press ctrl-c to abort.")
			continue
		}
		if stop {
			return false
		}

		if ok := b.Mark(sq, h.mark); !ok {
			fmt.Println("That square was already taken. Sorry.")

			av := make([]string, 0)
			for _, a := range b.Available() {
				av = append(av, fmt.Sprintf("%d", a))
			}
			fmt.Printf("Open squares are: %s\n", strings.Join(av, ", "))
			continue
		}

		return true
	}
}

// AskInput waits for user input.
// returns the square, a request to stop, or an error
func AskInput() (int, bool, error) {
	fmt.Println("Enter 0-8 to mark a square, or q to quit:")

	var input string
	count, err := fmt.Scan(&input)
	if err != nil {
		return 0, false, err
	}
	if count != 1 {
		return 0, false, fmt.Errorf("one can only mark a single square per turn")
	}
	if input == "q" {
		return 0, true, nil
	}
	sq, err := strconv.Atoi(input)
	if err != nil || sq < 0 || sq > 8 {
		return 0, false, fmt.Errorf("squares are indicated by the numbers 0 to 8")
	}

	return sq, false, nil
}
