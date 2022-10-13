package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	cur := []int{}
	line := ""
	for {
		cur = recurse(root,cur)
		if len(cur) == 0 {
			break
		}
	}
	return line
}



func recurse(root *TreeNode, cur []int ) []int {
	cur = append(cur, root.Val)
}