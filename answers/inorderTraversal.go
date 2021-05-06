package answers

//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
//94. 二叉树的中序遍历
func InorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return []int{}
	}

	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)          //左节点
		res = append(res, root.Val) //当前节点
		inorder(root.Right)         //右节点
	}

	inorder(root)
	return
}
