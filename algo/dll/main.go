package main

import (
	"fmt"
)

// doubly linked list

type Node struct {
	prop int
	prev *Node
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) AddToHead(prop int) {
	node := &Node{prop: prop}
	if list.head != nil {
		list.head.prev = node
		node.next = list.head
	}
	list.head = node
}

func (list *LinkedList) NodeWithValue(prop int) *Node {
	for node := list.head; node != nil; node = node.next {
		if node.prop == prop {
			return node
		}
	}
	return nil
}

func (list *LinkedList) AddAfter(prop int, node int) {
	new := &Node{prop: prop}
	nodeWith := list.NodeWithValue(node)
	if nodeWith != nil {
		new.prev = nodeWith
		new.next = nodeWith.prev
		nodeWith.next = new
	}

}

func (list *LinkedList) ItertateList() {
	for node := list.head; node != nil; node = node.next {
		fmt.Println(node.prop)
	}
}
func main() {
	list := &LinkedList{}
	list.AddToHead(1)
	list.AddAfter(2, 1)
	list.AddAfter(1, 2)
	list.ItertateList()

}
