package arraysandstrings

import (
	"fmt"
	"testing"
)

type palindromePermutationTestCase struct {
	str    string
	expect bool
}

var palindromePermutationTestCases = []palindromePermutationTestCase{
	{"", true},
	{"a", true},
	{"aa", true},
	{"aba", true},
	{"aaa", true},
	{"aabb", true},
	{"ab", false},
}

func TestPalindromePermutation(t *testing.T) {
	fmt.Println("Testing palindrome permutation function:")
	for index, testcase := range palindromePermutationTestCases {
		output := PalindromePermutation(testcase.str)
		if output != testcase.expect {
			t.Errorf("Case %d failed (string = %s)", index, testcase.str)
		}
	}
}
