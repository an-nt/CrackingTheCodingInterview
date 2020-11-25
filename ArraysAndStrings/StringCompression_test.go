package arraysandstrings

import "testing"

type stringCompressionTestCase struct {
	input  string
	expect string
}

var stringCompressionTestCases = []stringCompressionTestCase{
	{"", ""},
	{"a", "a"},
	{"aa", "aa"},
	{"aaa", "a3"},
	{"aaab", "aaab"},
	{"aaaab", "a4b1"},
	{"aaabb", "a3b2"},
	{"aaabbccc", "a3b2c3"},
	{"aaabbcccd", "a3b2c3d1"},
}

func TestStringCompression(t *testing.T) {
	for index, testcase := range stringCompressionTestCases {
		output := StringCompression(testcase.input)
		if output != testcase.expect {
			t.Errorf("Case %d failed: want %s but return %s", index, testcase.expect, output)
		}
	}
}
