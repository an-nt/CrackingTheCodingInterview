package silicon

import "testing"

type arraySumTestCase struct {
	input  []int
	expect int
}

var arraySumTestCases = []arraySumTestCase{
	{[]int{}, 0},
	{[]int{-1}, -1},
	{[]int{-1, 0, 3}, 2},
	{[]int{1, 2, 3, 4, 10, 11}, 31},
}

func Test_arraySum(t *testing.T) {
	for _, testcase := range arraySumTestCases {
		output := arraySum(testcase.input)
		if output != testcase.expect {
			t.Errorf("Sum of %d must be %d, but return %d", testcase.input, testcase.expect, output)
		}
	}
}
