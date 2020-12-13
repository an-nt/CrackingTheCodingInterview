package recursion

import "testing"

type canSumTestCase struct {
	targetSum int
	nums      []int
	expect    bool
}

var canSumTestCases = []canSumTestCase{
	{-1, []int{1, 2, 0}, false},
	{0, []int{5, 4, 7}, true},
	{8, []int{1, 4, 5}, true},
	{7, []int{3, 5, 9}, false},
	{1000000, []int{0, 1, 2, 3, 25}, true},
}

func Test_CanSum_memo(t *testing.T) {
	for _, testcase := range canSumTestCases {
		output := CanSum_memo(testcase.targetSum, testcase.nums)
		if output != testcase.expect {
			t.Errorf("The following case failed: sum %d nums %d", testcase.targetSum, testcase.nums)
		}
	}
}

func Test_CanSum_tabu(t *testing.T) {
	for _, testcase := range canSumTestCases {
		output := CanSum_tabu(testcase.targetSum, testcase.nums)
		if output != testcase.expect {
			t.Errorf("The following case failed: sum %d nums %d", testcase.targetSum, testcase.nums)
		}
	}
}
