package arraysandstrings

import "testing"

type oneAwayTestCase struct {
	str1   string
	str2   string
	expect bool
}

var oneAwayTestCases = []oneAwayTestCase{
	{"", "", true},
	{"", "a", true},
	{"b", "", true},
	{"a", "a", true},
	{"a", "b", true},
	{"b", "a", true},
	{"a", "ab", true},
	{"ab", "b", true},
	{"ab", "cb", true},
	{"ab", "ab", true},
	{"ab", "bc", false},
	{"ab", "cd", false},
	{"pale", "ple", true},
	{"pales", "pale", true},
	{"pale", "bale", true},
	{"pale", "bake", false},
}

func TestOneAway(t *testing.T) {
	for index, testcase := range oneAwayTestCases {
		output := OneAway(testcase.str1, testcase.str2)
		if output != testcase.expect {
			t.Errorf("Case %d (string %s and string %s) failed", index, testcase.str1, testcase.str2)
		}
	}
}
