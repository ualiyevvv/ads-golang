package binarytree

func (t *TreeNode) Add(val int) {
	if t == nil {
		return
	}

	if val > t.Val {
		if t.Right == nil {
			t.Right = &TreeNode{Val: val, Parrent: t}
		} else {
			t.Right.Add(val)
		}
	} else if val < t.Val {
		if t.Left == nil {
			t.Left = &TreeNode{Val: val, Parrent: t}
		} else {
			t.Left.Add(val)
		}
	} else {
		return
	}
}
