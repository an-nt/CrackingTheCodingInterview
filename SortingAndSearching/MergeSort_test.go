package SortingAndSearching

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	for i, testcase := range ArrayTestCases {
		output := MergeSort(testcase)

		if !IsAscendOrder(output) {
			t.Errorf("Test case No.%d failed: output = %d", i, output)
		}
	}
}
