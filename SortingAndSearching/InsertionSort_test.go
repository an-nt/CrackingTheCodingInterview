package SortingAndSearching

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	for i, testcase := range ArrayTestCases {
		InsertionSort(testcase)

		if !IsAscendOrder(testcase) {
			t.Errorf("Test case No.%d failed: output = %d", i, testcase)
		}
	}
}
