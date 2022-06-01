package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	for i := 0; btreeapplybylevel(root, f, i); i++ {
	}
}

func btreeapplybylevel(root *TreeNode, f func(...interface{}) (int, error), level int) bool {
	if root == nil {
		return false
	}
	if level == 0 {
		f(root.Data)
		return true
	}
	left := btreeapplybylevel(root.Left, f, level-1)
	right := btreeapplybylevel(root.Right, f, level-1)
	return left || right
}
