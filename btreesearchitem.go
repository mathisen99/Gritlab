package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}

	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}
	if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	}
	return root
}
