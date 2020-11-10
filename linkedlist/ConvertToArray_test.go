package linkedlist

import "testing"

func TestConvertLinkedListToArray(t *testing.T) {
	testcases := []struct {
		input  *SinglyLinkedList
		expect []int
	}{
		{&SinglyLinkedList{nil}, []int{}},
		{&SinglyLinkedList{&Node{4, nil}}, []int{4}},
		{&SinglyLinkedList{&Node{4, &Node{5, nil}}}, []int{4, 5}},
		{&SinglyLinkedList{&Node{4, &Node{5, &Node{-7, nil}}}}, []int{4, 5, -7}},
	}

	for index, testcase := range testcases {
		output := ConvertLinkedListToArray(testcase.input)
		if len(output) != len(testcase.expect) {
			t.Fatalf("Case failed: expect %d elements but return %d elements", len(testcase.expect), len(output))
		}
		for i := 0; i < len(testcase.expect); i++ {
			if testcase.expect[i] != output[i] {
				t.Errorf("Case %d failed at %dth element: want %d but return %d", index, i, testcase.expect[i], output[i])
			}
		}
	}
}
