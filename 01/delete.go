package main

// Delete removes the Item with value from the tree
func (t *TreeNode) Delete(value int) {
	t.remove(value)
}

func (tr *TreeNode) remove(value int) *TreeNode {

	if tr == nil {
		return nil
	}

	if value < tr.value {
		tr.left = tr.left.remove(value)
		return tr
	}
	if value > tr.value {
		tr.right = tr.right.remove(value)
		return tr
	}

	if tr.left == nil && tr.right == nil {
		tr = nil
		return nil
	}

	if tr.left == nil {
		tr = tr.right
		return tr
	}
	if tr.right == nil {
		tr = tr.left
		return tr
	}

	smallestValOnRight := tr.right
	for {
		//find smallest value on the right side
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}

	tr.value = smallestValOnRight.value
	tr.right = tr.right.remove(tr.value)
	return tr
}
