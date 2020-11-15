package sortingandsearching

import (
	arrstr "CodingInterview/arraysandstrings"
)

func GroupAnagram(arr []string) []string {
	if len(arr) < 2 {
		return arr
	}

	marker := make(map[string][]int)

	for i := 0; i < len(arr); i++ {
		marker[arrstr.SortStringBubble(arr[i])] = append(marker[arrstr.SortStringBubble(arr[i])], i)
	}

	result := []string{}
	for _, indexlist := range marker {
		for _, index := range indexlist {
			result = append(result, arr[index])
		}
	}
	return result
}
