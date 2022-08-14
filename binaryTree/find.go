package binarytree

func (t *TreeNode) FindMax() int {
	if t.Right != nil {
		return t.Right.FindMax()
	}
	return t.Val
}

func (t *TreeNode) FindMin() int {
	if t.Left != nil {
		return t.Left.FindMin()
	}
	return t.Val
}
