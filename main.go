package main

import (
	bt "ads-golang/binaryTree"
	"fmt"
)

func main() {
	tree := &bt.TreeNode{Val: 8}
	tree.Add(4)
	tree.Add(2)
	tree.Add(77)
	tree.Add(3)
	tree.Add(10)
	tree.Add(6)
	tree.Add(15)
	tree.Add(7)

	tree.InorderTraversal()
	fmt.Println("\nmax = ", tree.FindMax())
	fmt.Println("min = ", tree.FindMin())
	// tree.Delete(20)
	// tree.InorderTraversal()

	tree.Delete(77)
	fmt.Println()
	tree.InorderTraversal()
	fmt.Println("\nmax = ", tree.FindMax())
}