package linkedlist

import (
	"testing"
)

func TestRemoveDups(t *testing.T) {
	testcases := []struct {
		input  *SinglyLinkedList
		expect *SinglyLinkedList
	}{
		{
			input:  &SinglyLinkedList{},
			expect: &SinglyLinkedList{},
		},
		{
			input:  &SinglyLinkedList{&Node{0, nil}},
			expect: &SinglyLinkedList{&Node{0, nil}},
		},
		{
			input:  &SinglyLinkedList{&Node{5, &Node{-2, nil}}},
			expect: &SinglyLinkedList{&Node{5, &Node{-2, nil}}},
		},
		{
			input:  &SinglyLinkedList{&Node{7, &Node{7, nil}}},
			expect: &SinglyLinkedList{&Node{7, nil}},
		},
		{
			input:  &SinglyLinkedList{&Node{-7, &Node{0, &Node{5, nil}}}},
			expect: &SinglyLinkedList{&Node{-7, &Node{0, &Node{5, nil}}}},
		},
		{
			input:  &SinglyLinkedList{&Node{-7, &Node{0, &Node{-7, nil}}}},
			expect: &SinglyLinkedList{&Node{-7, &Node{0, nil}}},
		},
		{
			input:  &SinglyLinkedList{&Node{1, &Node{2, &Node{2, nil}}}},
			expect: &SinglyLinkedList{&Node{1, &Node{2, nil}}},
		},
		{
			input:  &SinglyLinkedList{&Node{1, &Node{1, &Node{2, nil}}}},
			expect: &SinglyLinkedList{&Node{1, &Node{2, nil}}},
		},
		{
			input:  &SinglyLinkedList{&Node{0, &Node{0, &Node{0, nil}}}},
			expect: &SinglyLinkedList{&Node{0, nil}},
		},
	}

	for index, testcase := range testcases {
		RemoveDups(testcase.input)
		outputNode := testcase.input.Head
		expectNode := testcase.expect.Head
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
