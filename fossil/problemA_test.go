package fossil

import "testing"

type problemATestCase struct {
	input  int
	expect int
}

var problemATestCases = []problemATestCase{
	{11, 10},
	{117, 99},
	{1, -1},
	{2, 1},
	{87, 75},
	{1001, 982},
	{1000000000, 999999932},
}

func TestProblemA(t *testing.T) {
	for index, testcase := range problemATestCases {
		output := ProblemA(testcase.input)
		if output != testcase.expect {
			t.Errorf("Case %d failed: want %d but return %d", index, testcase.expect, output)
		}
	}
}
