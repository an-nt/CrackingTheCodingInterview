package linkedlist

func RemoveDups(list SinglyLinkedList) {
	if list.Head.Next == nil {
		return
	}

	var existingValues map[int]bool
	existingValues[list.Head.Value] = true

	for node := list.Head; node.Next != nil; node = node.Next {
		_, existed := existingValues[node.Next.Value]
		if !existed {
			existingValues[node.Next.Value] = true
		} else {
			node.Next = node.Next.Next
		}
	}
}
