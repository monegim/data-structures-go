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

func (linkedList *LinkedList) IterateList() {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println("Property: ", node.property)
	}
}

func (linkedList *LinkedList) LastNode() *Node {
	var lastNode *Node
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{
		property: property,
		nextNode: nil,
	}
	lastNode := linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	} else {
		// List is empty
		linkedList.headNode = node
	}
}

func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			return node
		}
	}
	return nil
}

func (linkedList *LinkedList) AddAfter(nodeProperty, property int) {
	node := &Node{
		property: property,
		nextNode: nil,
	}
	nodeWith := linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	//fmt.Printf("LinkedList Head Property: %#v\n", linkedList.headNode.property)

	//linkedList.IterateList()
	//
	linkedList.AddToEnd(-1)
	//fmt.Println("Last Node: ", linkedList.LastNode().property)
	linkedList.IterateList()
	fmt.Println("the node with the, ", -1, "property is: ", linkedList.NodeWithValue(-1))
	linkedList.AddAfter(3, 5)
	linkedList.IterateList()
}
