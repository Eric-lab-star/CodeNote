package main

import "fmt"

type Node struct {
	key   int
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Search(value int) bool {
	return search(t.root, value)
}

func search(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if node.value < value {
		return search(node.right, value)
	}

	if node.value > value {
		return search(node.left, value)
	}

	return true
}

func (t *Tree) Insert(key, value int) {
	node := &Node{key, value, nil, nil}
	if t.root == nil {
		t.root = node
	}

	if !t.Search(value) {
		insert(t.root, node)
	}
}

func insert(root, node *Node) {
	if root.value < node.value {
		if root.right == nil {
			root.right = node
		} else {
			insert(root.right, node)
		}
	}
	if root.value > node.value {
		if root.left == nil {
			root.left = node
		} else {
			insert(root.left, node)
		}
	}
}

func (t *Tree) PrintTree() {
	print(t.root, 0)
}

func print(node *Node, level int) {
	if node != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "    "
		}
		level++
		format += ">> "
		print(node.right, level)
		fmt.Printf(format+"%d\n", node.value)
		print(node.left, level)

	}
}

func (t *Tree) Inorder() {
	inorder(t.root)
}

func inorder(root *Node) {
	if root != nil {
		inorder(root.left)
		fmt.Println(root.value)
		inorder(root.right)

	}
}

func (t *Tree) Preorder() {
	preorder(t.root)
}

func preorder(node *Node) {
	if node != nil {
		fmt.Println(node.value)
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
		fmt.Println(node.value)
	}
}

func (t *Tree) MinNode() *int {
	node := t.root
	if node == nil {
		return nil
	}
	for {
		if node.left == nil {
			return &node.value
		}
		node = node.left
	}
}

func (t *Tree) MaxNode() *int {
	treeNode := t.root
	if treeNode == nil {
		return nil
	}
	for {
		if treeNode.right == nil {
			return &treeNode.value
		}
		treeNode = treeNode.right
	}
}

func (t *Tree) RemoveNode(value int) {
	removeNode(t.root, value)
}


//removeNode!!!!
func removeNode(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if value < node.value {
		node.left = removeNode(node.left, value)
		return node
	}

	if value > node.value {
		node.right = removeNode(node.right, value)
		return node
	}

	//node without child

	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
			

	//node with 1 child
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}
	leftmostrightNode := node.right

	//node with 2 children

	for {
		if leftmostrightNode != nil && leftmostrightNode.left != nil {
			leftmostrightNode = leftmostrightNode.left
		} else {
			break
		}
	}

	//copy to node
	node.key, node.value = leftmostrightNode.key, leftmostrightNode.value
	//remove leftmostrightNode
	node.right = removeNode(node.right, node.key)
	return node
}

func main() {
	tree := &Tree{}
	tree.Insert(4, 4)
	tree.Insert(10, 10)
	tree.Insert(10, 10)
	tree.Insert(3, 3)
	n := *tree.MaxNode()
	fmt.Println(n)
	tree.PrintTree()
	tree.Inorder()
}
