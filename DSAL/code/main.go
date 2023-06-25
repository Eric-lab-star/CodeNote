package main

import "fmt"


type LinkedList struct{
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(prop int){
	node := &Node{property: prop}
	linkedList.headNode = node
	fmt.Printf("%d is added to head\n", linkedList.headNode.property)
}

func (linkedList *LinkedList) Iterate(){
	for node:= linkedList.headNode; node != nil; node = node.nextNode{
		fmt.Printf("current node is %d\n", node.property)
	}
}

func (linkedList *LinkedList) AddToEnd(prop int){
	
}

type Node struct{
	property int
	nextNode *Node
}



func main(){
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.Iterate()
}
