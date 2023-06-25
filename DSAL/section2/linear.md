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
