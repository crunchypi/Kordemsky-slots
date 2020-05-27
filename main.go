package main

import (
	"fmt"
	"os"

	arg "./arg-tools"
	boards "./board-tools"
	moves "./move-tools"
)

func main() {
	pass := arg.Check(os.Args[1:])
	if !pass {
		return
	}

	// Basic setup.
	board, err := boards.CreateBoard(os.Args[1:])
	if err != nil {
		fmt.Println("Create fail")
		return
	}
	goal := boards.ReverseSlice(board)

	// Trackers for movement.
	direction := true
	totalMoves := 0
	prevSteps := make([]int, 0, 10e2)

	for !boards.CheckEqual(board, goal) {
		// Keep track of previous to prevent backtracking.
		indexOfZero, _ := moves.IndexOfZero(board)
		prevSteps = append(prevSteps, indexOfZero)

		// Try step (long before short).
		moved := moves.MoveZero(board, false, direction, prevSteps)
		if !moved {
			moves.MoveZero(board, true, direction, prevSteps)
		}

		// # Turn direction on wall.
		direction = moves.HandleTurn(board, direction)
		// # For final return.
		totalMoves++

		fmt.Println(board)
	}

	fmt.Println("Done, total moves:", totalMoves)
}
