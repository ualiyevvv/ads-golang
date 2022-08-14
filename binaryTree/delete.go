package binarytree

func (t *TreeNode) Delete(val int) {
	if t == nil {
		return
	}

	if val < t.Val {
		t.Left.Delete(val)
	} else if val > t.Val {
		t.Right.Delete(val)
	} else {
		if t.Left == nil {
			if t.Parrent.Left == t {
				t.Parrent.Left = t.Right
			} else {
				t.Parrent.Right = t.Right
			}
		} else if t.Right == nil {
			if t.Parrent.Left == t {
				t.Parrent.Left = t.Left
			} else {
				t.Parrent.Right = t.Left
			}
		} else if t.Left != nil && t.Right != nil {
			t.Val = t.Right.FindMin()
			t.Right.Delete(t.Val)
		}

	}

}

// func (t *TreeNode) Search(val int) TreeNode {
// 	if t.Val == val {
// 		return *t
// 	}

// 	if val < t.Val {
// 		return t.Left.Search(val)
// 	} else if val > t.Val {
// 		return t.Right.Search(val)
// 	} else {
// 		defer fmt.Errorf("tree has not this value")
// 	}
// 	return *t
// }

// func (t *TreeNode) transplant(root, node *TreeNode) {
// 	if t == nil {
// 		return
// 	}

// 	if root == root.Right {
// 		root.Right = node
// 		return
// 	}
// 	root.Left = node
// }
