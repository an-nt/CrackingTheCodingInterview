package SortingAndSearching

import "testing"

func TestIsAscendOrder(t *testing.T) {
	testcases := []struct {
		input  []int
		expect bool
	}{
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, 2}, true},
		{[]int{1, 2, 2}, true},
		{[]int{1, 1, 1}, true},
		{[]int{-1, 0, 1}, true},
		{[]int{3, 2, 1}, false},
		{[]int{2, 2, 1}, false},
		{[]int{1, 0, -1}, false},
		{[]int{1, 2, 1}, false},
		{[]int{1, 2, 2, 1}, false},
		{[]int{1, 2, 3, 1}, false},
	}
	for _, testcase := range testcases {
		output := IsAscendOrder(testcase.input)
		if output != testcase.expect {
			t.Errorf("Test case failed with the input: %d\n", testcase.input)
		}
	}
}
