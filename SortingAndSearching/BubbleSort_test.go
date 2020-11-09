package sortingandsearching

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	for i, testcase := range ArrayTestCases {
		BubbleSort(testcase)

		if !IsAscendOrder(testcase) {
			t.Errorf("Test case No.%d failed: output = %d", i, testcase)
		}
	}
}
