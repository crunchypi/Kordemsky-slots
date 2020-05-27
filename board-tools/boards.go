package boards

import (
	"errors"
	"strconv"
)

// CheckEqual checks if two int slices are equal.
func CheckEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// ReverseSlice reverses the order of a slice.
// Not in-place.
func ReverseSlice(old []int) []int {
	new := make([]int, len(old))

	j := len(old) - 1
	for i := 0; i < len(old); i++ {
		new[i] = old[j]
		j--
	}
	return new
}

// CreateBoard ... converts a slice of strings into
// a 'gameboard' for this puzzle. Input is expected to
// by a slice of os.args
func CreateBoard(b []string) ([]int, error) {
	// # Attempt to convert str slice to int slice.
	new := make([]int, len(b), len(b))
	for i, v := range b {
		intVal, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			var empty []int
			return empty, errors.New("found non-int value in args")
		}
		new[i] = int(intVal)
	}
	return new, nil
}
