package linkedlist

import (
	"fmt"
	"strconv"
)

// Singly-linked lists are already defined with this interface:
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func AddTwoHugeNumbers(a *ListNode, b *ListNode) *ListNode {
	arrA := listToArray(a)
	arrB := listToArray(b)

	i := 0
	j := 0
	additional := 0
	arrResult := []int{}
	for i < len(arrA) && j < len(arrB) {
		arrResult = append(arrResult, (arrA[i]+arrB[j]+additional)%10000)
		if arrA[i]+arrB[j]+additional > 10000 {
			additional = 1
		} else {
			additional = 0
		}
	}
	for i < len(arrA) {
		arrResult = append(arrResult, (arrA[i]+additional)%10000)
		if arrA[i]+additional > 10000 {
			additional = 1
		} else {
			additional = 0
		}
		i++
	}
	for j < len(arrB) {
		arrResult = append(arrResult, (arrB[j]+additional)%10000)
		if arrB[j]+additional > 10000 {
			additional = 1
		} else {
			additional = 0
		}
		j++
	}
	if additional == 1 {
		arrResult = append(arrResult, 1)
	}
	return arrayToList(arrResult)
}

func listToArray(list *ListNode) []int {
	result := []int{}
	for node := list; node != nil; node = node.Next {
		nodeValue, _ := strconv.Atoi(fmt.Sprint(node.Value))
		result = append(result, nodeValue)
	}

	return result
}
func arrayToList(arr []int) *ListNode {
	result := &ListNode{}

	for i := 0; i < len(arr); i++ {
		node := &ListNode{
			Value: arr[i],
		}
		if result != nil {
			node.Next = result
		}
		result = node
	}
	return result
}
