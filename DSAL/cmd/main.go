package main

import "fmt"

//binary search tree

type Node struct {
	left  *Node
	right *Node
	key   int
	count int
}

type Tree struct {
	root *Node
}

func (t *Tree) Search(key int) *Node {
	return search(t.root, key)

}

func search(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if root.key < key {
		return search(root.right, key)
	}
	if root.key > key {
		return search(root.left, key)
	}
	return root
}

func (t *Tree) Insert(key int) {
	exist := t.Search(key)
	if exist == nil {
		node := &Node{nil, nil, key, 1}
		if t.root == nil {
			t.root = node
		} else {
			insert(t.root, node)
		}
	} else {
		exist.count++
	}
}

func insert(root, n *Node) {
	if root.key < n.key {
		if root.right == nil {
			root.right = n
		} else {
			insert(root.right, n)
		}
	} else {
		if root.left == nil {
			root.left = n
		} else {
			insert(root.left, n)
		}
	}

}

func (t *Tree) Delete(key int) {
	delete(t.root, key)
}

func delete(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key < key {
		node.right = delete(node.right, key)
		return node
	}
	if node.key > key {
		node.left = delete(node.left, key)
		return node
	}

	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}

	if node.left == nil {
		node = node.right
		return node
	}

	if node.right == nil {
		node = node.left
		return node
	}

	for leftmostRightNode := node.right; node != nil; leftmostRightNode = leftmostRightNode.left {
		if leftmostRightNode.left == nil {
			node.key = leftmostRightNode.key
			node.right = delete(node.right, leftmostRightNode.key)
			break
		}
	}

	return node

}

func (t *Tree) Inorder() {
	inorder(t.root)
}

func inorder(node *Node) {
	if node != nil {
		inorder(node.left)
		fmt.Println(node.key)
		inorder(node.right)
	}
}

func (t *Tree) Preorder() {
	preorder(t.root)
}

func preorder(node *Node) {
	if node != nil {
		fmt.Println(node.key)
		preorder(node.left)
		preorder(node.right)
	}
}

func (t *Tree) Postorder() {
	postorder(t.root)
}

func postorder(node *Node) {
	if node != nil {
		postorder(node.left)
		postorder(node.right)
		fmt.Println(node.key)

	}
}

func height(node *Node) int {

	if node == nil {
		return 0
	}
	return max(height(node.left), height(node.right)) + 1
}
func col(h int) int {
	if h == 1 {
		return 1
	}
	return col(h-1) + col(h-1) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (t *Tree) Print() {
	fmt.Println("---------------")
	fmt.Println("---------------")
}
func pow(a int, b int) int {
	for i := 0; i < b; i++ {
		a = a * a
	}
	return a
}
func print(m *[][]int, n *Node, col int, height int, row int) {
	if n == nil {
		return
	}
	print(m, n.left, col-pow(2, height-2), row+1, height-1)
	print(m, n.left, col+pow(2, height-2), row+1, height-1)
}

func main() {
	tree := &Tree{}
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(8)
	tree.Insert(10)
	tree.Insert(1)
	tree.Insert(3)
	tree.Print()
	h := height(tree.root)

	if h != 0 {
		for i := 0; i < h; i++ {

		}
	}
}
