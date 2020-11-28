package linkedlist

import "testing"

type add2HugeNumTestCase struct {
	a      *ListNode
	b      *ListNode
	expect *ListNode
}

var add2HugeNumTestCases = []add2HugeNumTestCase{
	{
		a:      &ListNode{},
		b:      &ListNode{},
		expect: &ListNode{},
	},
}

func TestAddTwoHugeNumbers(t *testing.T) {
	for index, testcase := range add2HugeNumTestCases {
		output := AddTwoHugeNumbers(testcase.a, testcase.b)
		expectNode := testcase.expect
		outputNode := output
		for expectNode != nil || outputNode != nil {
			if expectNode.Value != outputNode.Value {
				t.Errorf("Case %d failed: wrong value", index)
			}
			expectNode = expectNode.Next
			outputNode = outputNode.Next
		}
		if (expectNode == nil) == (outputNode == nil) {
			t.Errorf("Case %d failed: wrong length", index)
		}
	}
}
