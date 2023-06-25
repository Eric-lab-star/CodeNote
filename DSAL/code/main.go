package main

import "fmt"

type Node struct{
	property int
	nextNode *Node
}

type LinkedList struct{
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int){
	node := Node{}
	node.property = property
	if node.nextNode != nil{
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node
}
func (linkedList *LinkedList) IterateList(){
	for node := linkedList.headNode; node != nil; node = node.nextNode{
		fmt.Println(node.property)
	}
}

func (linkedList *LinkedList) LastNode() *Node{
	lastNode := &Node{}
	for node := linkedList.headNode; node != nil; node = node.nextNode{
		if node.nextNode == nil{
			lastNode = node
		}
	}
	return lastNode
}


func (linkedList *LinkedList)AddToEnd(property int){
	node := &Node{property: property, nextNode: nil}
	lastNode :=	linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}
func (linkedList *LinkedList) NodeWithValue(property int) *Node{
	nodeWith := &Node{}
	for node := linkedList.headNode; node!= nil; node = node.nextNode{
		if node.property == property{
			nodeWith = node
			return nodeWith
		}
	}
	return nodeWith
	
}
func (linkedList *LinkedList) AddAfter(nodeProperty int, property int){
	node := &Node{
		property: property,
	}
	nodeWith := linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil{
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}
func main(){
	linkedList := LinkedList{}		
	linkedList.AddToHead(1)
	linkedList.AddToEnd(3)
	linkedList.AddAfter(1,8)
	linkedList.IterateList()
}


