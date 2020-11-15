package sortingandsearching

import (
	"testing"
)

func TestShellSort(t *testing.T) {
	for i, testcase := range ArrayTestCases {
		ShellSort(testcase)

		if !IsAscendOrder(testcase) {
			t.Errorf("Test case No.%d failed: output = %d", i, testcase)
		}
	}
}
