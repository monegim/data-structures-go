package main

type Node_2[T any] struct {
	property     T
	nextNode     *Node_2[T]
	previousNode *Node_2[T]
}

type DoublyLinkedList struct {
	headNode *Node
}
