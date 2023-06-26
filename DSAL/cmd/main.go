package main

import "fmt"

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(prop int) {
	node := &Node{prop, nil}
	linkedList.headNode = node
}

func (linkedList *LinkedList) IterateList() {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func (linkedList *LinkedList) LastNode() *Node {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			return node
		}
	}
	return &Node{}
}

func (linkedList *LinkedList) AddToEnd(prop int) {
	node := linkedList.LastNode()
	if node != nil {
		node.nextNode = &Node{prop, nil}
	}
}

func (linkedList *LinkedList) NodeWithValue(prop int) *Node {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == prop {
			return node
		}
	}
	return nil
}

func (linkedList *LinkedList) AddAfter(nodeProp, prop int) {
	node := &Node{prop, nil}
	prevNode := linkedList.NodeWithValue(nodeProp)
	node.nextNode = prevNode.nextNode
	prevNode.nextNode = node
}

type Node struct {
	property int
	nextNode *Node
}

func main() {
	linkedList := &LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToEnd(2)
	linkedList.AddToEnd(3)
	linkedList.AddAfter(2, 10)

	linkedList.IterateList()
}
