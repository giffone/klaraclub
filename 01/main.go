package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func (tr *TreeNode) findMax() int {
	if tr.right == nil {
		return tr.value
	}
	return tr.right.findMax()
}

func (tr *TreeNode) findMin() int {
	if tr.left == nil {
		return tr.value
	}
	return tr.left.findMin()
}

func (tr *TreeNode) delete(value int) {
	if tr == nil {
		return
	}
	if tr.value == value {
		if tr.left == nil && tr.right == nil {
			tr = nil
			return
		}
		if tr.left != nil && tr.right == nil {
			tr.value = tr.left.value
			tr.left = nil
			return
		}
		if tr.right != nil && tr.left == nil {
			tr.value = tr.right.value
			tr.right = nil
			return
		}
		min := tr.right.findMin()
		tr.delete(min)
		tr.value = min
	}
	tr.left.delete(value)
	tr.right.delete(value)
}

func (tr *TreeNode) insert(value int) error {
	if tr == nil {
		return nil
	}
	if value == tr.value {
		return errors.New("value is exist")
	}
	if value < tr.value {
		if tr.left == nil {
			tr.left = &TreeNode{value: value}
		} else {
			tr.left.insert(value)
		}
	} else {
		if tr.right == nil {
			tr.right = &TreeNode{value: value}
		} else {
			tr.right.insert(value)
		}
	}
	return nil
}

func main() {
	// tr := &TreeNode{value: 8}

	// for _, v := range []int{4, 2, 3, 10, 6, 7} {
	// 	tr.insert(v)
	// }
	// fmt.Println("max is ", tr.findMax())
	// fmt.Println("min is ", tr.findMin())
	// tr.printInOrder()
	// fmt.Println("\n-------")
	// tr.delete(4)
	// tr.printInOrder()

	tr2 := &TreeNode{value: 4}
	for _, v := range []int{2, 1, 3, 6, 5, 7} {
		tr2.insert(v)
	}
	list := tr2.dfs(3)
	for _, t := range list {
		fmt.Printf("%d -> ", t.value)
	}
}

func (tr *TreeNode) printInOrder() {
	if tr == nil {
		return
	}
	tr.left.printInOrder()
	fmt.Printf("%d -> ", tr.value)
	tr.right.printInOrder()
}
