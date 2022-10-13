package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) (sum int) {
	if root == nil {
		return 0
	}
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			sum += root.Left.Val
		}
		sum += sumOfLeftLeaves(root.Left)
	}

	if root.Right != nil {
		sum += sumOfLeftLeaves(root.Right)
	}
	return sum
}
