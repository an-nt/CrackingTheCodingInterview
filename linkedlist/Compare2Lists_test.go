package linkedlist

import "testing"

type compare2ListsTestCase struct {
	a      *SinglyLinkedList
	b      *SinglyLinkedList
	expect bool
}

var compare2ListsTestCases = []compare2ListsTestCase{
	{
		a:      &SinglyLinkedList{nil},
		b:      &SinglyLinkedList{nil},
		expect: true,
	},
	{
		a:      &SinglyLinkedList{&Node{1, &Node{0, &Node{-4, &Node{1, nil}}}}},
		b:      &SinglyLinkedList{&Node{1, &Node{0, &Node{-4, &Node{1, nil}}}}},
		expect: true,
	},
	{
		a:      &SinglyLinkedList{&Node{1, &Node{0, &Node{-4, nil}}}},
		b:      &SinglyLinkedList{&Node{1, &Node{0, &Node{-4, nil}}}},
		expect: true,
	},
	{
		a:      &SinglyLinkedList{&Node{1, nil}},
		b:      &SinglyLinkedList{&Node{1, &Node{0, &Node{-4, nil}}}},
		expect: false,
	},
	{
		a:      &SinglyLinkedList{&Node{4, nil}},
		b:      &SinglyLinkedList{nil},
		expect: false,
	},
}

func TestCompare2Lists(t *testing.T) {
	for index, testcase := range compare2ListsTestCases {
		output := Compare2Lists(testcase.a, testcase.b)
		if output != testcase.expect {
			t.Errorf("Case %d failed", index)
		}
	}
}
