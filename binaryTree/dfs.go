package binarytree

func (t *TreeNode) Dfs() []*TreeNode {
	visited := []*TreeNode{}
	if t == nil {
		return visited
	}

	return t.recurse(visited)
}

func (t *TreeNode) recurse(visited []*TreeNode) []*TreeNode {
	visited = append(visited, t)
	if t.Left != nil {
		visited = t.Left.recurse(visited)
	}

	if t.Right != nil {
		visited = t.Right.recurse(visited)
	}

	return visited
}
