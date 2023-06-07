# Data Structure and Algorithms

## Classification of data Structure and Structural design patterns

- Linear
  lists/ sets/ tuples/ queues/ stacks/ heap
- Non-Liner
  Trees/ tables/ containers
- Homogenous
  2d arrays/ MultiDArray
- Heterogenous
  Linked Lists/ Ordered Lists/ Unordered Lists
- Dynamic
  Dictionaries/ TreeSets/ Sequences

## List, Tuples and Heaps

### Lists

A lists is sequence of elements.
Each element can be connected to another with a link.
List have a variable length and developer can remove or add elements more easily

### Tuples

Tuples are typically immutable sequential collections.

### Heaps

A heap is a data structure that is based on the heap property.
The heap data structure is used in selection, graph, and k-way merge algorithims.
Operations such as finding, merging, insertion, key changes and deleting are performed on heaps.

if the order is descending, it is referred to as maximum heap; otherwise, it's a minimumheap.
It is paritally ordered data structure.

## Structural Design Patterns

### Adapter

**adapter wraps one of the objects to hide the complexity of conversion happening behinde the scenes.**

### Bridge

Abstraction is a high-level control layer for somw entity. This layer isn't supposed to do any real work on its own. It should delegate the work to the implementation layer

### Composite

composit is a structural design pattern that lets you compose objects into tree structures and then work with these structures as if they were individual objects.

A composite is a group of similar objects in a sinlge objects. Objects are stored in a tree form to persis the hirearchy. New types of objects can be added without changing the interface and the clinet code.