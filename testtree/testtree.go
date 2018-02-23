package main

import (
	"fmt"

	"github.com/OrthantBing/bst"
)

func main() {
	var bsearchtree bst.LinkedBinarySearchTree

	bsearchtree.AddElementToTree(10)
	bsearchtree.AddElementToTree(20)
	bsearchtree.AddElementToTree(15)

	// var linkedbinarytree linkedbinarytree.LinkedBinaryTree
	// rootnode := linkedbinarytree.Add_root(10)
	// leftnode := linkedbinarytree.Add_left(rootnode, 5)
	// rightnode := linkedbinarytree.Add_right(rootnode, 15)

	// linkedbinarytree.Add_left(leftnode, 2)
	// linkedbinarytree.Add_right(leftnode, 7)

	// linkedbinarytree.Add_left(rightnode, 8)
	// linkedbinarytree.Add_right(rightnode, 18)

	// leftleftnode := linkedbinarytree.Add_left(leftnode, 2)
	// rightleftnode := linkedbinarytree.Add_right(leftnode, 7)

	// leftrightnode := linkedbinarytree.Add_left(rightnode, 8)
	// rightrightnode := linkedbinarytree.Add_right(rightnode, 18)
	fmt.Println(bsearchtree)
	//fmt.Println(rootnode)

}
