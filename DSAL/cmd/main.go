package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
	prevNode *Node
}
type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) NodeBetweenValues(firstProperty, secondProperty int) *Node {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode != nil && node.prevNode != nil {
			if node.nextNode.property == secondProperty && node.prevNode.property == firstProperty {
				return node
			}
		}
	}
	return nil
}

func (linkedList *LinkedList) AddToHead(prop int) {
	node := &Node{property: prop}
	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
		linkedList.headNode.prevNode = node
	}
	linkedList.headNode = node
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
	node := &Node{property: prop}
	nodeWith := linkedList.NodeWithValue(nodeProp)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.prevNode = nodeWith
		nodeWith.nextNode = node
	}
}

func (linkedList *LinkedList) LastNode() *Node {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			return node
		}
	}
	return nil
}

func (linkedList *LinkedList) AddToEnd(prop int) {
	node := &Node{prop, nil, nil}
	lastNode := linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.prevNode = lastNode
	}
}

func (linkedList *LinkedList) IterateList() {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func main() {
	linkedList := &LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToEnd(2)
	linkedList.AddToEnd(3)
	node := linkedList.NodeBetweenValues(1, 3)
	fmt.Println(node)
	linkedList.IterateList()
}
