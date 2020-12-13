package others

import "testing"

type trimSpaceTestCase struct {
	input  string
	expect string
}

var trimSpaceTestCases = []trimSpaceTestCase{
	{" ", ""},
	{" a", "a"},
	{" a ", "a"},
	{"  ab   ", "ab"},
	{" a b  c  ", "a b c"},
	{" a  hello  c  ", "a hello c"},
}

func TestTrimSpace(t *testing.T) {
	for index, testcase := range trimSpaceTestCases {
		output := TrimSpace(testcase.input)
		if output != testcase.expect {
			t.Errorf("Case %d failed: want %s but return %s", index, testcase.expect, output)
		}
	}
}
