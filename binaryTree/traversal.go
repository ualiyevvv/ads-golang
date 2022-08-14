package binarytree

import "fmt"

func (t *TreeNode) InorderTraversal() {
	if t == nil {
		return
	}

	t.Left.InorderTraversal()
	fmt.Print(t.Val, " ")
	t.Right.InorderTraversal()
}
