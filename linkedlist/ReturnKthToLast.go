package linkedlist

var index int

//ReturnKthToLast return the Kth-to-last node
//it uses recursive algorithm to search the node
func ReturnKthToLast(list *SinglyLinkedList, k int) *Node {
	if list.Head == nil || k < 1 {
		return nil
	}
	if list.Head.Next == nil {
		index = 1
		if index == k {
			return list.Head
		}
		return nil
	}
	node := ReturnKthToLast(&SinglyLinkedList{Head: list.Head.Next}, k)
	if node != nil {
		return node
	}
	index++
	if index == k {
		return list.Head
	}
	return nil
}
