package linkedbinarytree

import (
	"fmt"
)

type Node struct {
	Element int
	Parent  *Node
	Left    *Node
	Right   *Node
}

func (n Node) String() string {
	return fmt.Sprintf(
		"Element: %v \n Left: %q \n Right: %q",
		n.Element, n.Left, n.Right)
}

func (lbt *Node) Num_children() int {
	var children int = 0
	if lbt.Left != nil {
		children += 1
	}
	if lbt.Right != nil {
		children += 1
	}
	return children
}

type LinkedBinaryTree struct {
	root *Node
	size int
}

func (lbt LinkedBinaryTree) String() string {
	return fmt.Sprintf(
		"Root: %q, Size: %v",
		lbt.root, lbt.size)
}

func (lbt *LinkedBinaryTree) Root() *Node {
	return lbt.root
}

func (lbt *LinkedBinaryTree) IsRoot(node *Node) bool {
	if lbt.root == node {
		return true
	}
	return false
}

func (lbt *LinkedBinaryTree) Size() int {
	return lbt.size
}

func (lbt *LinkedBinaryTree) Add_root(e int) *Node {
	if lbt.root != nil {
		fmt.Println("Root exists")
		return lbt.Root()
	}
	lbt.size++
	lbt.root = &Node{e, nil, nil, nil}
	return lbt.Root()
}

func (lbt *LinkedBinaryTree) Add_left(node *Node, e int) *Node {
	if node.Left != nil {
		fmt.Println("Left exists")
		return node.Left
	}
	node.Left = &Node{e, nil, nil, nil}
	//node.Left.Parent = node
	return node.Left
}

func (lbt *LinkedBinaryTree) Add_right(node *Node, e int) *Node {
	if node.Right != nil {
		fmt.Println("Right exists")
		return node.Right
	}
	node.Right = &Node{e, nil, nil, nil}
	node.Right.Parent = node
	return node.Right
}

func (lbt *LinkedBinaryTree) Replace(node *Node, e int) int {
	oldElement := node.Element
	node.Element = e
	return oldElement
}

func (lbt *LinkedBinaryTree) Delete(node *Node) *Node {
	if node.Num_children() == 2 {
		fmt.Println("Cant delete node")
		return nil
	}
	var child *Node
	if node.Left == nil {
		child = node.Right
	} else {
		child = node.Left
	}

	if child != nil {
		child.Parent = node.Parent
	}
	if lbt.IsRoot(node) {
		lbt.root = child
	} else {
		Parent := node.Parent
		if node == Parent.Left {
			Parent.Left = child
		} else {
			Parent.Right = child
		}
		lbt.size--

	}
	return node
}
