package arg

import (
	"fmt"
	"strconv"
)

// Check verifies that most game-rules are followed.
// Returned bool represents fail/pass (True is pass).
func Check(args []string) bool {

	// # Guard minimum argument count.
	if len(args) < 3 {
		printHelp(1)
		return false
	}

	// # Guard int-only.
	for _, v := range args {
		_, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			printHelp(2)
			return false
		}
	}

	// # Guard permitted board-slots.
	for _, v := range args {
		if !(v == "0" || v == "1" || v == "2") {
			printHelp(3)
			return false
		}
	}

	// # Check quantities.
	counts := map[string]int{
		"0": 0,
		"1": 0,
		"2": 0,
	}
	for _, v := range args {
		counts[v] = counts[v] + 1
	}
	if counts["0"] != 1 {
		printHelp(4)
		return false
	}

	if counts["1"] != counts["2"] {
		printHelp(5)
		return false
	}

	// # Pass.
	return true
}

func printHelp(scenario int) {
	switch scenario {
	case 1:
		fmt.Println("Not enough arguments.")
	case 2:
		fmt.Println("A non-integer argument was found.")
	case 3:
		fmt.Println("Only allowed board pieces are: 0,1,2")
	case 4:
		fmt.Println("Zero must be listed once.")
	case 5:
		fmt.Println("Quantities of 1 and 2 are not the same.")
	default:
		fmt.Println("Uncaught error")
	}
	fmt.Println(
		"\nRules: This is a very specific puzzle.",
		"\n(1) The board length has to be an odd number, only possible pieces are: 0,1,2.",
		"\n(2) Goal is to swap all pieces represented by 1 with 2. Any piece can either move one or two tiles each turn.",
		"\n(3) There must be an equal amount of 1s and 2s, but only one 0, which represents an empty slot.",
		"\n(4) Pieces should be contigious, though this is not enforced (might not find a solution).",
		"\nEXAMPLE arguments: 1 1 0 2 2")

}
