package silicon

import "testing"

type leftRotationTestCase struct {
	array  []int
	d      int
	expect []int
}

var leftRotationTestCases = []leftRotationTestCase{
	{
		array:  []int{},
		d:      0,
		expect: []int{},
	},
	{
		array:  []int{1},
		d:      -1,
		expect: []int{1},
	},
	{
		array:  []int{1},
		d:      1,
		expect: []int{1},
	},
	{
		array:  []int{2},
		d:      2,
		expect: []int{2},
	},
	{
		array:  []int{1, 2},
		d:      3,
		expect: []int{1, 2},
	},
	{
		array:  []int{1, 2},
		d:      1,
		expect: []int{2, 1},
	},
	{
		array:  []int{1, 2, 3},
		d:      2,
		expect: []int{3, 1, 2},
	},
	{
		array:  []int{1, 2, 3, 4, 5},
		d:      4,
		expect: []int{5, 1, 2, 3, 4},
	},
	{
		array:  []int{1, 2, 3, 4, 5},
		d:      5,
		expect: []int{1, 2, 3, 4, 5},
	},
}

func Test_leftRotation(t *testing.T) {
	for _, testcase := range leftRotationTestCases {
		output := leftRotation(testcase.array, testcase.d)
		if len(output) != len(testcase.expect) {
			t.Errorf("Testcase %d failed: wrong output array length", testcase.array)
		} else {
			for j := 0; j < len(output); j++ {
				if output[j] != testcase.expect[j] {
					t.Errorf("Testcase %d failed: wrong value at index %d", testcase.array, j)
				}
			}
		}
	}
}
