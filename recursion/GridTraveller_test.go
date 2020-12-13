package recursion

import "testing"

type gridTravellerTestCase struct {
	row    int
	col    int
	expect int
}

var gridTravellerTestCases = []gridTravellerTestCase{
	{1, 1, 1},
	{2, 3, 3},
	{0, 5, -1},
	{1, -1, -1},
	{3, 3, 6},
	{18, 18, 2333606220},
	{20, 20, 35345263800},
}

func Test_GridTraveller(t *testing.T) {
	for i, testcase := range gridTravellerTestCases {
		output := GridTraveller(testcase.row, testcase.col)
		if output != testcase.expect {
			t.Errorf("Case %d failed: want %d but return %d", i, testcase.expect, output)
		}
	}
}
