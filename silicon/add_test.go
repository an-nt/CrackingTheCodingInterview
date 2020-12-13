package silicon

import "testing"

type addTestCase struct {
	input  int
	expect int
}

var addTestCases = []addTestCase{
	{-1, 0},
	{0, 0},
	{1, 1},
	{2, 3},
	{3, 6},
	{4, 10},
}

func Test_add(t *testing.T) {
	for _, testcase := range addTestCases {
		output := add(testcase.input)
		if output != testcase.expect {
			t.Errorf("add(%d) must be equal to %d, but return %d", testcase.input, testcase.expect, output)
		}
	}
}
