package bst

import (
	"fmt"

	"github.com/OrthantBing/linkedbinarytree"
)

type LinkedBinarySearchTree struct {
	linkedbinarytree.LinkedBinaryTree
}

func (lbt *LinkedBinarySearchTree) Subtree_search(n *linkedbinarytree.Node, val int) *linkedbinarytree.Node {
	if n.Element == val {
		return n
	} else if val < n.Element {
		if n.Left != nil {
			return lbt.Subtree_search(n.Left, val)
		}

	} else if val > n.Element {
		if n.Right != nil {
			return lbt.Subtree_search(n.Right, val)
		}
	}
	return n
}

func (lbt *LinkedBinarySearchTree) subtree_first_position(n *linkedbinarytree.Node) *linkedbinarytree.Node {
	walk := n
	for walk.Left != nil {
		walk = walk.Left
	}
	return walk
}

func (lbt *LinkedBinarySearchTree) subtree_last_position(n *linkedbinarytree.Node) *linkedbinarytree.Node {
	walk := n
	for walk.Right != nil {
		walk = walk.Right
	}
	return walk
}

func (lbt *LinkedBinarySearchTree) First() *linkedbinarytree.Node {
	if lbt.Size() == 0 {
		return nil
	}
	return lbt.subtree_first_position(lbt.Root())
}

func (lbt *LinkedBinarySearchTree) Last() *linkedbinarytree.Node {
	if lbt.Size() == 0 {
		return nil
	}
	return lbt.subtree_last_position(lbt.Root())
}

func (lbt *LinkedBinarySearchTree) GetItem(e int) *linkedbinarytree.Node {
	if lbt.Size() == 0 {
		fmt.Println("Element doesnt exist")
		return nil
	}
	node := lbt.Subtree_search(lbt.Root(), e)

	if node.Element != e {
		fmt.Println("Element doesnt exist")
		return nil
	}
	return node
}

func (lbt *LinkedBinarySearchTree) AddElementToTree(e int) {
	if lbt.Size() == 0 {
		lbt.Add_root(e)
	} else {
		p := lbt.Subtree_search(lbt.Root(), e)
		if p.Element > e {
			lbt.Add_left(p, e)
			return
		} else if p.Element < e {
			lbt.Add_right(p, e)
			return
		} else {
			p.Element = e
		}
	}

}
