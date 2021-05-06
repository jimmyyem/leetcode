package answers

//https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
//107. 二叉树的层序遍历
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func LevelOrderBottom(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	var depth float64
	q := make([][]*TreeNode, 0)
	temp := make([]*TreeNode, 0)
	for {
		if depth == 0 {
			q[depth] = append(q[depth], []*TreeNode{root.Left, root.Right})
			temp = append(temp, root.Left, root.Right)
		} else {

		}
		depth++

		//设置结束条件，到最后一层没有内容
		if len(q) == 0 {
			break
		}
	}

	return
}
