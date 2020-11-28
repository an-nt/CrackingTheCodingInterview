package linkedlist

func Compare2Lists(a *SinglyLinkedList, b *SinglyLinkedList) bool {
	nodeA := a.Head
	nodeB := b.Head
	for nodeA != nil && nodeB != nil {
		if nodeA.Value != nodeB.Value {
			return false
		}
		nodeA = nodeA.Next
		nodeB = nodeB.Next
	}
	if (nodeA == nil) != (nodeB == nil) {
		return false
	}
	return true
}
