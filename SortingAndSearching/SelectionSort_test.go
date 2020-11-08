package SortingAndSearching

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	for i, testcase := range ArrayTestCases {
		SelectionSort(testcase)

		if !IsAscendOrder(testcase) {
			t.Errorf("Test case No.%d failed: output = %d", i, testcase)
		}
	}
}
