package main

import "fmt"

type Node struct {
	prop     int
	nextNode *Node
}

type List struct {
	head *Node
}

func (list *List) AdddToHead(prop int) {
	node := &Node{prop: prop}
	if list.head != nil {
		node.nextNode = list.head
	}
	list.head = node
}

func (list *List) NodeWithValue(prop int) *Node {
	for node := list.head; node != nil; node = node.nextNode {
		if node.prop == prop {
			return node
		}
	}
	return nil
}

func (list *List) AddAfter(target, prop int) {
	node := &Node{prop: prop}
	nodeWith := list.NodeWithValue(target)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}
func (list *List) LastNode() *Node {
	for node := list.head; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			return node
		}
	}
	return nil

}

func (list *List) AddToEnd(prop int) {
	node := &Node{prop: prop}
	lastNode := list.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

func (list *List) IterateList() {
	for node := list.head; node != nil; node = node.nextNode {
		fmt.Println(node)
	}

}
func main() {
	list := &List{}
	list.AdddToHead(1)
	list.AddAfter(1, 2)
	list.AddToEnd(4)
	list.IterateList()
}
