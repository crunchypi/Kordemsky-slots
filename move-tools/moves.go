package moves

import "errors"

// MoveZero attempts to move the only 'zero' in this specific puzzle,
// according to game rules.
func MoveZero(board []int, short, direction bool, prevSteps []int) bool {
	indexOfZero, _ := IndexOfZero(board)
	var target int
	var step int
	// # Assess how far zero should step.
	if short {
		step = 1
	} else {
		step = 2
	}
	// # Assess direction zero should go.
	if direction {
		target = indexOfZero + step
	} else {
		target = indexOfZero - step
	}
	// # Prevent backtracking.
	if len(prevSteps) > 1 {
		if prevSteps[len(prevSteps)-2] == target {
			return false
		}
	}
	// # Prevent out-of-bounds.
	if target < 0 || target > len(board)-1 {
		return false
	}
	// # Move.
	board[indexOfZero] = board[target]
	board[target] = 0
	return true
}

// HandleTurn checks if a direction should be reversed.
// This function returns the opposite of 'currentDirection'
// if the index of zero is out-of-bounds.
func HandleTurn(board []int, currentDirection bool) bool {
	indexOfZero, _ := IndexOfZero(board)
	if indexOfZero == 0 || indexOfZero == len(board)-1 {
		return !currentDirection
	}
	return currentDirection
}

// IndexOfZero checks which current index zero
// is at in the board.
func IndexOfZero(board []int) (int, error) {
	for i, v := range board {
		if v == 0 {
			return i, nil
		}
	}
	return 0, errors.New("no zero")
}
