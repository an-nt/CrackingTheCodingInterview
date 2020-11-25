package arraysandstrings

import "testing"

func TestIsPermutation(t *testing.T) {
	testCases := []struct {
		str1   string
		str2   string
		expect bool
	}{
		{"abc", "bca", true},
		{"abcht", "bca", false},
		{"god", "God", false},
	}

	for _, testCase := range testCases {
		output := IsPermutation(testCase.str1, testCase.str2)
		if output != testCase.expect {
			t.Errorf("Want %t, Have %t", testCase.expect, output)
		}
	}
}

func TestSortString(t *testing.T) {
	testcases := []struct {
		input  string
		expect string
	}{
		{"cba", "abc"},
		{"hat", "aht"},
		{"hatkx", "ahktx"},
		{"hgfedcba", "abcdefgh"},
	}

	for _, testcase := range testcases {
		output := SortStringBubble(testcase.input)
		if output != testcase.expect {
			t.Errorf("Want %s, Have %s", testcase.expect, output)
		}
	}
}
