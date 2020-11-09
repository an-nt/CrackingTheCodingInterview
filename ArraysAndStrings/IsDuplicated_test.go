package arraysandstrings

import (
	"testing"
)

func TestIsDuplicated(t *testing.T) {
	testcases := []struct {
		input  []int
		expect bool
	}{
		{[]int{}, false},
		{[]int{-1}, false},
		{[]int{-1, 2}, false},
		{[]int{-1, 2, 2}, true},
		{[]int{0, -1, 2, -1}, true},
	}
	for _, testcase := range testcases {
		output := IsDuplicated(testcase.input)
		if output != testcase.expect {
			t.Errorf("Case failed: %d", testcase.input)
		}
	}
}
