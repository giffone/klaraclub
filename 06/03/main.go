package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	allVal := []int{}
	if root == nil {
		return false
	}
	allVal = recurse(root, allVal)

	for i := 0; i < len(allVal); i++ {
		for j := 1; j < len(allVal); j++ {
			if allVal[i] + allVal[j] == k && i != j {
				return true
			}
		}
	}
	return false
}

func recurse(root *TreeNode, allVal []int) []int {
	allVal = append(allVal, root.Val)
	if root.Left != nil {
		allVal = recurse(root.Left, allVal)
	}
	if root.Right != nil {
		allVal = recurse(root.Right, allVal)
	}
	return allVal
}