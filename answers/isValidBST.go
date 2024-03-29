package answers

import "math"

//https://leetcode-cn.com/problems/validate-binary-search-tree/
//98. 验证二叉搜索树
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func IsValidBST(root *TreeNode) bool {
	return validBST(root, math.MinInt64, math.MaxInt64)

	//nodeList := preOrderTree(root)
	//
	////判断是否升序
	//for i := 1; i < len(nodeList); i++ {
	//	if nodeList[i-1] >= nodeList[i] {
	//		return false
	//	}
	//}
	//
	//return true
}

//前序遍历二叉树
func preOrderTree(root *TreeNode) (res []int) {
	if root == nil {
		return []int{}
	}

	var preOrder func(*TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		preOrder(root.Left)         //左节点
		res = append(res, root.Val) //当前节点
		preOrder(root.Right)        //右节点
	}

	preOrder(root)
	return
}

/**
minVal
maxVal
*/
func validBST(root *TreeNode, minVal, maxVal int) bool {
	// 这是代表到最后的子节点了
	if root == nil {
		return true
	}

	// 所有判断是否合规的条件只有这一条，针对每个节点传入的值不同
	if root.Val <= minVal || root.Val >= maxVal {
		return false
	}

	// 分别判断左节点和右节点，然后递归
	return validBST(root.Left, minVal, root.Val) &&
		validBST(root.Right, root.Val, maxVal)
}
