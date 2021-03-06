package linkedlist

import (
	"testing"
)

type removeDupsTestCase struct {
	input  *SinglyLinkedList
	expect *SinglyLinkedList
}

var removeDupsTestCases = []removeDupsTestCase{
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
	{
		input:  &SinglyLinkedList{&Node{-7, &Node{0, &Node{5, &Node{8, nil}}}}},
		expect: &SinglyLinkedList{&Node{-7, &Node{0, &Node{5, &Node{8, nil}}}}},
	},
	{
		input:  &SinglyLinkedList{&Node{-7, &Node{0, &Node{-4, &Node{-4, nil}}}}},
		expect: &SinglyLinkedList{&Node{-7, &Node{0, &Node{-4, nil}}}},
	},
	{
		input:  &SinglyLinkedList{&Node{1, &Node{0, &Node{-4, &Node{1, nil}}}}},
		expect: &SinglyLinkedList{&Node{1, &Node{0, &Node{-4, nil}}}},
	},
	{
		input:  &SinglyLinkedList{&Node{-7, &Node{-7, &Node{5, &Node{5, nil}}}}},
		expect: &SinglyLinkedList{&Node{-7, &Node{5, nil}}},
	},
	{
		input:  &SinglyLinkedList{&Node{-7, &Node{0, &Node{-7, &Node{0, nil}}}}},
		expect: &SinglyLinkedList{&Node{-7, &Node{0, nil}}},
	},
	{
		input:  &SinglyLinkedList{&Node{6, &Node{0, &Node{0, &Node{6, nil}}}}},
		expect: &SinglyLinkedList{&Node{6, &Node{0, nil}}},
	},
	{
		input:  &SinglyLinkedList{&Node{4, &Node{4, &Node{4, &Node{4, nil}}}}},
		expect: &SinglyLinkedList{&Node{4, nil}},
	},
}

func TestRemoveDups(t *testing.T) {
	for index, testcase := range removeDupsTestCases {
		RemoveDups(testcase.input)
		passed := Compare2Lists(testcase.input, testcase.expect)
		if !passed {
			t.Errorf("Case %d failed", index)
		}
	}
}
