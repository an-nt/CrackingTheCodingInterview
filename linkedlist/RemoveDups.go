package linkedlist

func RemoveDups(list *SinglyLinkedList) {
	if list.Head == nil {
		return
	}
	if list.Head.Next == nil {
		return
	}

	//existingValues stores value in the list
	existingValues := make(map[int]bool)

	//runner technique: use two pointers
	//the "node" pointer moves earlier than the "previous" pointer by one step
	previous := &Node{}

	for node := list.Head; node != nil; node = node.Next {
		_, existed := existingValues[node.Value]
		if !existed {
			existingValues[node.Value] = true
			previous = node
		} else {
			previous.Next = node.Next
		}
	}
}
