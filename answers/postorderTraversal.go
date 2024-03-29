package answers

//https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
//145. 二叉树的后序遍历（左节点，右节点，当前节点）
func PostorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return []int{}
	}

	var postorder func(*TreeNode)
	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		postorder(root.Left)        //左节点
		postorder(root.Right)       //右节点
		res = append(res, root.Val) //当前节点
	}

	postorder(root)
	return
}

func postorderTraversal(root *TreeNode) (res []int) {
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

		if node.Right != nil {
			stack.Push(node.Right)
		}

		res = append(res, node.Val)
	}

	return
}
