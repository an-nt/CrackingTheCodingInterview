package others

import "testing"

type deleteWithMinCostTestCase struct {
	str    string
	cost   []int
	expect int
}

var deleteWithMinCostTestCases = []deleteWithMinCostTestCase{
	{
		str:    "",
		cost:   []int{},
		expect: 0,
	},
	{
		str:    "a",
		cost:   []int{1},
		expect: 0,
	},
	{
		str:    "aa",
		cost:   []int{1, 2},
		expect: 1,
	},
	{
		str:    "aaa",
		cost:   []int{1, 2, 3},
		expect: 3,
	},
	{
		str:    "abaa",
		cost:   []int{1, 2, 3, 2},
		expect: 2,
	},
	{
		str:    "abccbd",
		cost:   []int{0, 1, 2, 3, 4, 5},
		expect: 2,
	},
	{
		str:    "aabbcc",
		cost:   []int{1, 2, 1, 2, 1, 2},
		expect: 3,
	},
	{
		str:    "aaaa",
		cost:   []int{3, 4, 5, 6},
		expect: 12,
	},
	{
		str:    "ababa",
		cost:   []int{10, 5, 10, 5, 10},
		expect: 0,
	},
}

func TestDeleteWithMinCost(t *testing.T) {
	for index, testcase := range deleteWithMinCostTestCases {
		output := DeleteWithMinCost(testcase.str, testcase.cost)
		if output != testcase.expect {
			t.Errorf("Case %d failed: want %d but return %d", index, testcase.expect, output)
		}
	}
}
