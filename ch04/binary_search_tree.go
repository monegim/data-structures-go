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
func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	inOrderTraverseTree(tree.rootNode, function)
}

func inOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		inOrderTraverseTree(treeNode.leftNode, function)
		function(treeNode.value)
		inOrderTraverseTree(treeNode.rightNode, function)
	}
}

func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	preOrderTraverseTree(tree.rootNode, function)
}

//  preOrderTraverseTree method
func preOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		function(treeNode.value)
		preOrderTraverseTree(treeNode.leftNode, function)
		preOrderTraverseTree(treeNode.rightNode, function)
	}
}

// PostOrderTraverseTree method
func (tree *BinarySearchTree) PostOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	postOrderTraverseTree(tree.rootNode, function)
}

//  postOrderTraverseTree method
func postOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		postOrderTraverseTree(treeNode.leftNode, function)
		postOrderTraverseTree(treeNode.rightNode, function)
		function(treeNode.value)
	}
}

func main() {
	rootNode := &TreeNode{
		key:   15,
		value: 15,
	}
	bst := BinarySearchTree{
		rootNode: rootNode,
	}
	bst.Insert(6, 6)
	bst.Insert(23, 23)
	bst.Insert(71, 71)
	bst.Insert(50, 50)
	bst.Insert(4, 4)
	bst.Insert(7, 7)
	bst.Insert(5, 5)
	fmt.Println("Order")
	bst.InOrderTraverseTree(func(i int) { fmt.Println(i) })
	fmt.Println("PreOrder")
	bst.PreOrderTraverseTree(func(i int) { fmt.Println(i) })
	fmt.Println("PostOrder")
	bst.PostOrderTraverseTree(func(i int) { fmt.Println(i) })

}
