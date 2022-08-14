package binarytree

func (t *TreeNode) Dfs(val int) []*TreeNode {
	visited := []*TreeNode{}
	if t == nil {
		return visited
	}

	return t.recurse(visited, val)
}

var flag = false

func (t *TreeNode) recurse(visited []*TreeNode, val int) []*TreeNode {
	if flag {
		return visited
	}
	visited = append(visited, t)

	if t.Val == val || flag {
		flag = true
		return visited
	}

	if t.Left != nil {
		visited = t.Left.recurse(visited, val)
	}
	if t.Right != nil {
		visited = t.Right.recurse(visited, val)
	}

	return visited

}
