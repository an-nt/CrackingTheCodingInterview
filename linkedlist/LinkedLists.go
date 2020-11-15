package linkedlist

type SinglyLinkedList struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}

type ISinglyLinkedList interface {
	RemoveNode(target *Node)
}
