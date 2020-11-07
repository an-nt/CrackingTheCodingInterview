package SortingAndSearching

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	for _, testcase := range ArrayTestCases {
		BubbleSort(testcase)
		if !IsAscendOrder(testcase) {
			t.Errorf("Test case failed at case: %d", testcase)
		}
	}
}
