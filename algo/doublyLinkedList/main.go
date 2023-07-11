package main

import "fmt"

type Node struct {
	prop     int
	nextNode *Node
	prevNode *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) AddToHead(prop int) {
	node := &Node{prop: prop}
	if list.head != nil {
		node.nextNode = list.head
		list.head.prevNode = node
	}
	list.head = node
}

func (list *LinkedList) NodeWithValue(prop int) *Node {
	for node := list.head; node != nil; node = node.nextNode {
		if node.prop == prop {
			return node
		}
	}
	return nil
}

func (list *LinkedList) AddAfter(target, prop int) {
	nodeWith := list.NodeWithValue(target)
	node := &Node{prop: prop}

	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.prevNode = nodeWith
		nodeWith.nextNode.prevNode = node
		nodeWith.nextNode = node

	}
}

func (list *LinkedList) LastNode() *Node {
	for node := list.head; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			return node
		}
	}
	return nil
}

func (list *LinkedList) AddToEnd(prop int) {
	lastNode := list.LastNode()
	node := &Node{prop: prop}
	if lastNode != nil {
		lastNode.nextNode = node
		node.prevNode = lastNode
	}
}

func (list *LinkedList) IterateToNext() {
	for node := list.head; node != nil; node = node.nextNode {
		fmt.Println(node)
	}
}

func (list *LinkedList) IterateToPrev() {
	for node := list.LastNode(); node != nil; node = node.prevNode {
		fmt.Println(node)
	}
}

func (list *LinkedList) NodeBetween(prev, next int) *Node {
	for node := list.head; node != nil; node = node.nextNode {
		if node.prevNode != nil && node.nextNode != nil {
			if node.prevNode.prop == prev && node.nextNode.prop == next {
				return node
			}
		}
	}
	return nil
}

func main() {
	list := &LinkedList{}
	list.AddToHead(1)
	list.AddToEnd(2)
	list.AddToEnd(4)
	fmt.Println(list.NodeBetween(1, 4))

}
