package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	replace := node
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		replace.Parent.Left = rplc
	} else {
		replace.Parent.Right = rplc
	}
	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}
