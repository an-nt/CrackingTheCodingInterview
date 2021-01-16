package sortingandsearching

import "testing"

func TestHeapSort(t *testing.T) {
	for i, testcase := range ArrayTestCases {
		HeapSort(testcase)

		if !IsAscendOrder(testcase) {
			t.Errorf("Test case No.%d failed: output = %d", i, testcase)
		}
	}
}
