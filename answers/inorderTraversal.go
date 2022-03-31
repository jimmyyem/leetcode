package answers

//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
//94. 二叉树的中序遍历（左节点，当前节点，右节点）
func InorderTraversal(root *TreeNode) (res []int) {
	return inorderTraversal(root)

	//if root == nil {
	//	return []int{}
	//}
	//
	//var inorder func(*TreeNode)
	//inorder = func(root *TreeNode) {
	//	if root == nil {
	//		return
	//	}
	//
	//	inorder(root.Left)          //左节点
	//	res = append(res, root.Val) //当前节点
	//	inorder(root.Right)         //右节点
	//}
	//
	//inorder(root)
	//return
}

func inorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return []int{}
	}

	stack := NewStackNode()

	// 把所有Left入栈
	for root != nil {
		stack.Push(root)
		root = root.Left
	}

	for !stack.IsEmpty() {
		pop, _ := stack.Pop()
		node := pop.(*TreeNode)

		res = append(res, node.Val)

		if node.Right != nil {
			stack.Push(node.Right)
		}
	}

	return
}
