package piscine

func BTreeIsBinary(root *TreeNode) bool {
	left := true
	right := true
	if root == nil {
		return true
	}

	if root.Left != nil {
		left = BTreeIsBinary(root.Left) && root.Data >= root.Left.Data
	}

	if root.Right != nil {
		right = BTreeIsBinary(root.Right) && root.Data <= root.Right.Data
	}

	return left && right
}
