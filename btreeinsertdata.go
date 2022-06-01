package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	node := &TreeNode{Data: data}
	if root.Data < data {
		if root.Right == nil {
			root.Right = node
			root.Right.Parent = root
		} else {
			BTreeInsertData(root.Right, data)
		}
	} else {
		if root.Left == nil {
			root.Left = node
			root.Left.Parent = root
		} else {
			BTreeInsertData(root.Left, data)
		}
	}
	return root
}
