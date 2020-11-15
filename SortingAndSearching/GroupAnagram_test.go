package sortingandsearching

import (
	"testing"

	arrstr "CodingInterview/arraysandstrings"
)

type groupAnagramTestCase struct {
	input []string
}

var groupAnagramTestCases = []groupAnagramTestCase{
	{[]string{}},
	{[]string{"a"}},
	{[]string{"a", "b"}},
	{[]string{"ab", "cd", "ba", "dc"}},
	{[]string{"a", "abc", "cba", "cba", "cb", "bc"}},
	{[]string{"a!", "!a", "ab", "bc", "ca", "ab"}},
}

func TestGroupAnagram(t *testing.T) {
	for index, testcase := range groupAnagramTestCases {
		output := GroupAnagram(testcase.input)
		//just test cases that have more than two elements
		if len(output) >= 2 {
			//wasRangeStoped indicates whether the range of elements that have the same sorted string is stoped
			wasRangeStoped := make(map[string]bool)
			for i, value := range output {
				_, existed := wasRangeStoped[arrstr.SortStringBubble(value)]
				if !existed {
					wasRangeStoped[arrstr.SortStringBubble(value)] = false
					if i != 0 {
						wasRangeStoped[arrstr.SortStringBubble(output[i-1])] = true
					}
				} else {
					if wasRangeStoped[arrstr.SortStringBubble(value)] {
						t.Errorf("Case %d failed", index)
					}
				}
			}
		}

	}
}
