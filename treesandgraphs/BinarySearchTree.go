package treesandgraphs

import (
	"errors"
	"sync"
)

type BinarySearchTree struct {
	root *TreeNode
	lock sync.Mutex
}

type IBinarySearchTree interface {
	AddNodeToTree(key int, value int) error
	InOrderTraverse(function func(int))
	PreOrderTraverse(function func(int))
	PostOrderTraverse(function func(int))
	RemoveNode(key int)
}

func (b *BinarySearchTree) addNodeToTree(key int, value int) error {
	b.lock.Lock()
	defer b.lock.Unlock()
	node := &TreeNode{
		key:   key,
		value: value,
	}
	if b.root == nil {
		b.root = node
		return nil
	}
	return addNodeToTree(b.root, node)
}

func addNodeToTree(tree *TreeNode, node *TreeNode) error {
	if node.key < tree.key {
		if tree.left == nil {
			tree.left = node
			return nil
		}
		return addNodeToTree(tree.left, node)
	}
	if node.key > tree.key {
		if tree.right == nil {
			tree.right = node
			return nil
		}
		return addNodeToTree(tree.right, node)
	}
	return errors.New("Duplicated keys")
}

//InOrderTraverse follows the order: left -> node -> right
func (b *BinarySearchTree) InOrderTraverse(function func(int)) {
	b.lock.Lock()
	defer b.lock.Unlock()
	inOrderTraverse(b.root, function)
}

func inOrderTraverse(tree *TreeNode, function func(int)) {
	if tree == nil {
		return
	}
	inOrderTraverse(tree.left, function)
	function(tree.value)
	inOrderTraverse(tree.right, function)
}

//PreOrderTraverse follows the order: node -> left -> right
func (b *BinarySearchTree) PreOrderTraverse(function func(int)) {
	b.lock.Lock()
	defer b.lock.Unlock()
	preOrderTraverse(b.root, function)
}

func preOrderTraverse(tree *TreeNode, function func(int)) {
	if tree == nil {
		return
	}
	function(tree.value)
	preOrderTraverse(tree.left, function)
	preOrderTraverse(tree.right, function)
}

//PostOrderTraverse follows the order: left -> right -> node
func (b *BinarySearchTree) PostOrderTraverse(function func(int)) {
	b.lock.Lock()
	defer b.lock.Unlock()
	postOrderTraverse(b.root, function)
}

func postOrderTraverse(tree *TreeNode, function func(int)) {
	if tree == nil {
		return
	}
	postOrderTraverse(tree.left, function)
	postOrderTraverse(tree.right, function)
	function(tree.value)
}

func (b *BinarySearchTree) RemoveNode(key int) {
	b.lock.Lock()
	defer b.lock.Unlock()
	removeNode(b.root, key)
}

func removeNode(tree *TreeNode, key int) *TreeNode {
	if tree == nil {
		return tree
	}

	if tree.key < key {
		tree.right = removeNode(tree.right, key)
		return tree
	}

	if tree.key > key {
		tree.left = removeNode(tree.left, key)
		return tree
	}

	if tree.left == nil && tree.right == nil {
		tree = nil
		return tree
	}

	if tree.left != nil {
		tree = tree.left
		return tree
	}

	if tree.right != nil {
		tree = tree.right
		return tree
	}

	var leftMostRightNode *TreeNode
	for leftMostRightNode = tree.right; leftMostRightNode.left != nil; {
		leftMostRightNode = leftMostRightNode.left
	}
	tree.key, tree.value = leftMostRightNode.key, leftMostRightNode.value
	tree.right = removeNode(tree.right, key)
	return tree
}
