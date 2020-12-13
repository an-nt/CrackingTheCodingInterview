package fossil

import "testing"

type problemCTestCase struct {
	input  int
	expect string
}

var problemCTestCases = []problemCTestCase{
	{1, "Bob"},
	{2, "Bob"},
	{3, "Bob"},
	{4, "Alice"},
	{6, "Bob"},
	{17, "Bob"},
	{21, "Bob"},
	{55, "Bob"},
	{40, "Alice"},
	{68, "Alice"},
	{88, "Alice"},
	{90, "Bob"},
	{96, "Alice"},
}

func TestProblemC(t *testing.T) {
	for index, testcase := range problemCTestCases {
		output := ProblemC(testcase.input)
		if output != testcase.expect {
			t.Errorf("Case %d failed", index)
		}
	}
}
