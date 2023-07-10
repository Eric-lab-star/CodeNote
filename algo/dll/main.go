package main

import "fmt"

type node struct {
	prop     int
	nextNode *node
}

type list struct {
	head *node
}

func (list *list) AddToHead(prop int) {
	node := &node{prop: prop}
	if list.head != nil {
		node.nextNode = list.head
	}
	list.head = node
}

func (list *list) AddAfter(target, prop int) {
	node := &node{prop: prop}
	for targetNode := list.head; targetNode != nil; targetNode = targetNode.nextNode {
		if targetNode.prop == target {
			node.nextNode = targetNode.nextNode
			targetNode.nextNode = node
		}
	}
}

func (list *list) NodeWithValue(prop int) *node {
	for node := list.head; node != nil; node = node.nextNode {
		if node.prop == prop {
			return node
		}
	}
	return nil
}

func (list *list) LastNode() *node {
	for node := list.head; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			return node
		}
	}
	return nil
}

func (list *list) AddToLast(prop int) {
	node := &node{prop: prop}
	lastNode := list.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

func (list *list) Iterate() {
	for node := list.head; node != nil; node = node.nextNode {
		fmt.Println(node)
	}
}
func main() {
	list := &list{}
	list.AddToHead(1)
	list.AddToLast(3)
	list.AddToLast(4)
	list.Iterate()
}
