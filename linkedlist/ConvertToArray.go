package linkedlist

func ConvertLinkedListToArray(list *SinglyLinkedList) []int {
	result := make([]int, 0)
	if list.Head == nil {
		return result
	}

	for currentNode := list.Head; currentNode != nil; currentNode = currentNode.Next {
		result = append(result, currentNode.Value)
	}
	return result
}
