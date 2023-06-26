# Linear Data Structures

## List

List is a collection of ordered elements that are used to store list of items.
Unlike array lists, these can expand and shrink dynamically.

### LinkedList

LinkedList is a sequence of nodes that have properties and a refernece to the next node in the sequence.
They are not stored contiguously in memory.

#### Node Class

```go
type Node struct{
    property int
    nextNode *Node
}
```

#### LinkedList Class

```go
type LinkedList struct{
    headNode *Node
}
```

#### AddToHead method

```go
func (linkedList *LinkedList) AddToHead(property int){
    var node = Node{}
    node.property = property
    if node.nextNode != nil{
        node.nextNode = linkedList.headNode

    }
    linkedLists.headNode = &node
}
```

### Doubly linked List

In a doubly linked list, ass nodes have a pointer to the node they are connected to on either side of them. This means that each node is connected to two nodes, and we can traverse forward through to the next node or backward through to the previous node.

```go
type Node struct{
    property int
    nextNode *Node
    previousNode *Node
}
```

#### NodeBetweenValues Method

```go
func (linkedList *LinkedList) NodeBetweenValues(firstProperty, secondProperty int) *Node{
    for node := linkedList.headNode; node != nil; node = node.nextNode{
        if node.previousNode != nil && node.nextNode != nil{
            if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty{
                return node
            }
        }
    }
    return nil
}
```

#### AddToHead method

```go
func (linkedList *LinkedList) AddToHead(property int){
    node := &Node{
          property: property,
          }
    if linkedList.headNode != nil{
        node.nextNode = linkedList.headNode
        linkedList.headNode.previousNode = node
    }
    linkedList.headNode = node
}
```

#### AddAfter method

Addafter method adds a node after a specific node to a double linked list.

```go
func (linkedList *LinkedList ) AddAfter(nodeProperty, prop in){
    node := &Node{property: prop}
    nodeWith := linkedList.NodeWithValue(nodeProperty)
    if nodeWith != nil{
        node.nextNode = nodeWith.nextNode
        node.previousNode = nodeWith
        nodeWith.nextNode = node
    }
}
```


#### AddToEnd method

```go
func (linkedList *LinkedList) AddToend(property int){
    node := &Node{property, nil, nil}
    lastNode := linkedList.LastNode()
    if lastNode != nil{
        lastNode.nextNode = node
        node.previousNode = lastNode
    }
}
```
### Sets

A set is a linear data structure that has a collection of values that are not repeated. A set can store unique values without any particular order.

```go
type Set struct{
    integerMap map[int]bool
}

func (set *Set) New(){
    set.integerMap = make(map[int]bool)
}
```

#### AddElement 

```go
func (set *Set) AddElement(element int){
    if !set.ContainElement(element){
        set.integerMap[element] = true
    }
}
```

#### DeletElement method

```go
func (set *Set) DeleteElement(element int){
    delete(set.integerMap, element)
}
```

#### ContainsElement 

```go
func (set *Set) ContainsElement(element in)bool{
    _, exists := set.integerMap[element]
    return exists

```
#### InterSect method

```go
func (set *Set) Intersect(anotherSet *Set) *Set{
    intersectSet := &Set{}
    interSectSet.New()
    for v := range set.integerMap{
        if anotherSet.ContainElement(value){
            intersectSet.AddElement(value)
        }
    }
    return intersectSet

}
```


#### Union method


