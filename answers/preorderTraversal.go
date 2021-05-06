package answers

//https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
//144. 二叉树的前序遍历
func PreorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return []int{}
	}

	var preorder func(*TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val) //当前节点
		preorder(root.Left)         //左节点
		preorder(root.Right)        //右节点
	}

	preorder(root)
	return
}
