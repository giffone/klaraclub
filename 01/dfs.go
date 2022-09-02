package main

func (tr *TreeNode) dfs(value int) []*TreeNode {
	visited := []*TreeNode{}

	if tr == nil {
		return visited
	}

	return tr.recurse(value, visited)
}

func (tr *TreeNode) recurse(value int, visited []*TreeNode) []*TreeNode {
	visited = append(visited, tr)
	
	if tr.left != nil {
		visited = tr.left.recurse(value, visited)
		if tr.value == value {
			return visited
		}
	}
	if tr.right != nil {
		visited = tr.right.recurse(value, visited)
	}
	return visited
}
