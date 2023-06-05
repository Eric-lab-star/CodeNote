# Data Structures and Algorithms

Data Structure refers to the organization of the data in a computer memory, in orderto retrieve it quickly for processing.

## Classification

- Linear

  - Lists
  - Sets
  - Tuples
  - Queues
  - Stacks
  - Heap

- None-Linear

  - Trees
  - Tables
  - Containers

- Homogeneous

  - 2D Arrays
  - MultiDArrays

- Heterogeneous

  - Linked List
  - Ordered Lists
  - unordered Lists

- Dynamic

  - Dictionaries
  - Treesets
  - Sequences

## Lists

### Description

- Sequence of elements
- Each elements can be connected to another with a link in a forward or backward direction.
- The element can have other payload properties.
- Basice type of container.
- Elements need not be contiguous in memory or on disk.

## Tuples

### Description

- Finite soreted list of elements
- Immutable sequential collections
- related fields of different data types

## Heaps

### Description

- Data Structure bses on heap property
- Used in seletion, graph and k-way merge
- The value stored at each node is greater than or qual to its children, max heap only
- descending order is called maxheap; opposite is min heap
- It is partially ordered

## Adapter

The Adapter pattern provides a wrapper with a interface required by the API clinet to link incompatiable types nad act as a translator between two types.

The adapter uses the interface of a class to be a a class with another compatible interface.

- **dependency inversion** principle
- **Delegation principle**

Multiple formats handling source-to-destination transformations are the scenarios where the adapter petterns is applied.

The adapter pattern comprieses the target, adaptee, adapter, and client

- target is the interface that the client calls and invokes methods on the adapter and adaptee
- The client wants the incompatible interace implemented by the adapter.
- The adapter translates the incompatible interface of the adaptee into an interface that the client wants.

## Bridge

Bridge decouples the implementation from the abstraction(interface). The abstract base class(interface) can be subclassed to provide different implementations and allow implemntation details to be modified eaily.

The interface, which is a bridge, helps in making the functionality for concrete classes independent from the interface implementer classes.

The bridge pattern components are abstraction, refined abstraction, implementer, and concrete implementer.
Abstraction is the interface implemented as an abstract class that clients invoke with the method on the concrete implementer.

Abscraction maintains a has-a relationship with the implementation, instead of an is-a relationship. The has-a relationship is maintained by composition. Abstraction has a reference of the implementation. Refined abstraction provides more variations than abstraction.

Abstraction `IContour`
Refined Abstraction `DrawContour`
Implementation `IDrawShape`
concrete implementer `DrawShape`
