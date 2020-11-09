package sortingandsearching

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	for i, testcase := range ArrayTestCases {
		QuickSort(testcase, 0, len(testcase)-1)

		if !IsAscendOrder(testcase) {
			t.Errorf("Test case No.%d failed: output = %d", i, testcase)
		}
	}
}
