package linkedlist

import "testing"

type returnKthToLastTestCase struct {
	list   *SinglyLinkedList
	k      int
	expect *Node
}

var returnKthToLastTestCases = []returnKthToLastTestCase{
	{&SinglyLinkedList{}, 0, nil},
	{&SinglyLinkedList{}, -1, nil},
	{&SinglyLinkedList{}, 1, nil},
	{&SinglyLinkedList{&Node{7, nil}}, 0, nil},
	{&SinglyLinkedList{&Node{7, nil}}, -1, nil},
	{&SinglyLinkedList{&Node{7, nil}}, 1, &Node{7, nil}},
	{&SinglyLinkedList{&Node{7, &Node{-5, nil}}}, -1, nil},
	{&SinglyLinkedList{&Node{7, &Node{-5, nil}}}, 0, nil},
	{&SinglyLinkedList{&Node{7, &Node{-5, nil}}}, 1, &Node{-5, nil}},
	{&SinglyLinkedList{&Node{7, &Node{-5, nil}}}, 2, &Node{7, &Node{-5, nil}}},
	{&SinglyLinkedList{&Node{0, &Node{0, nil}}}, 1, &Node{0, nil}},
	{&SinglyLinkedList{&Node{0, &Node{0, nil}}}, 2, &Node{0, &Node{0, nil}}},
}

func TestReturnKthToLast(t *testing.T) {
	for index, testcase := range returnKthToLastTestCases {
		output := ReturnKthToLast(testcase.list, testcase.k)
		if !(output == nil && testcase.expect == nil) {
			if (output != nil) != (testcase.expect != nil) {
				t.Fatalf("Case %d failed: either output or expect is nil", index)
			}
			if (output != nil) == (testcase.expect != nil) {
				if output.Value != testcase.expect.Value {
					t.Errorf("Case %d failed: wrong value, expect %d but return %d", index, testcase.expect.Value, output.Value)
				}
				outputNode := output
				expectNode := testcase.expect
				for outputNode != nil || expectNode != nil {
					if outputNode.Value != expectNode.Value {
						t.Errorf("Case %d failed: wrong values", index)
					}
					outputNode = outputNode.Next
					expectNode = expectNode.Next
					if (outputNode == nil) != (expectNode == nil) {
						t.Fatalf("Case %d failed: wrong length", index)
						break
					}
				}
			}
		}

	}
}
