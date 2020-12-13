package recursion

import (
	"fmt"
)

// GridTraveller returns how many ways to go from the top-left to the bottom-right corner
// of an (row x col) matrix. Known that you can only move downwards or rightwards, 1 step each.
func GridTraveller(row int, col int) int {
	if row < 1 || col < 1 {
		return -1
	}

	memo := make(map[string]int)

	memo["1,1"] = 1
	return gridTraveller(row, col, memo)
}

// Complexity:
// Time: row*col
// Space: row^2 * col^2, because it calls distinct recursive functions (row*col) times
// each time we cost the space complexity of O(row*col) for the memo variable

func gridTraveller(row int, col int, memo map[string]int) int {
	if row == 0 || col == 0 {
		return 0
	}

	key := fmt.Sprint(row) + "," + fmt.Sprint(col)
	_, existed := memo[key]
	if existed {
		return memo[key]
	}

	memo[key] = gridTraveller(row-1, col, memo) + gridTraveller(row, col-1, memo)
	return memo[key]
}
