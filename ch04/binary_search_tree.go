package main

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	key, value          int
	leftNode, rightNode *TreeNode
}

type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

func (bst *BinarySearchTree) Insert(key, value int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	node := &TreeNode{
		key, value, nil, nil,
	}
	if bst.rootNode == nil {
		bst.rootNode = node
	} else {
		insertTreeNode(bst.rootNode, node)
	}
}

func insertTreeNode(rootNode, newTreeNode *TreeNode) {
	if newTreeNode.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newTreeNode
		} else {
			insertTreeNode(rootNode.leftNode, newTreeNode)
		}
	} else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = newTreeNode
		} else {
			insertTreeNode(rootNode.rightNode, newTreeNode)
		}
	}
}

// InOrderTraverseTree method
func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)) {}
func main() {
	rootNode := &TreeNode{
		key:   53,
		value: 100,
	}
	bst := BinarySearchTree{
		rootNode: rootNode,
	}
	bst.Insert(7, 100)
	fmt.Println(bst.rootNode)

}
