package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil
	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = node
}
func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	fmt.Printf("LinkedList Head Property: %#v", linkedList.headNode.property)
}
