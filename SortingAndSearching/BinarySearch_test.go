package sortingandsearching

import "testing"

func TestBinarySearch(t *testing.T) {
	testcases := []struct {
		sortedArray []int //the array must be sorted in an ascending order
		value       int
		expect      int
	}{
		//nil arrays
		{[]int{}, 1, -1},
		{[]int{}, -1, -1},
		//arrays with one element
		{[]int{1}, 1, 0},
		{[]int{0}, 1, -1},
		{[]int{-1}, 1, -1},
		{[]int{-1}, -1, 0},
		//arrays with two elements
		{[]int{1, 2}, 2, 1},
		{[]int{1, 3}, 0, -1},
		{[]int{1, 3}, 2, -1},
		{[]int{1, 3}, 4, -1},
		{[]int{1, 2}, -2, -1},
		{[]int{-1, -2}, 2, -1},
		{[]int{-1, -2}, -2, 1},
		//arrays with three elements
		{[]int{1, 2, 3}, 1, 0},
		{[]int{1, 2, 3}, 2, 1},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{2, 4, 6}, 1, -1},
		{[]int{2, 4, 6}, 3, -1},
		{[]int{2, 4, 6}, 5, -1},
		{[]int{2, 4, 6}, 7, -1},
		{[]int{1, 2, 3}, -1, -1},
		{[]int{-3, -2, -1}, 1, -1},
		{[]int{-3, -2, -1}, -3, 0},
		{[]int{-3, -2, -1}, -2, 1},
		{[]int{-3, -2, -1}, -1, 2},
		//long arrays
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -5, -1},
		{[]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4}, 0, 5},
		{[]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4}, 5, -1},
		{[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1}, 5, -1},
		{[]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1}, -5, 5},
	}

	for _, testcase := range testcases {
		output := BinarySearch(testcase.sortedArray, testcase.value)
		if output != testcase.expect {
			t.Errorf("Case failed: Search %d in %d, want %d but return %d", testcase.value, testcase.sortedArray, testcase.expect, output)
		}
	}
}
