package answers

//https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
//102. 二叉树的层序遍历
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func LevelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	var depth float64
	q := make([]*TreeNode, 0)
	for {
		if depth == 0 {
			res = append(res, []int{root.Val})
			q = append(q, root.Left, root.Right)
		} else {
			tmpSlice := make([]int, 0)
			size := len(q) //当前q长度
			for i := 0; i < size; i++ {
				if q[i] != nil {
					tmpSlice = append(tmpSlice, q[i].Val)
					q = append(q, q[i].Left, q[i].Right)
				}
			}
			//本层有内容才放入res
			if len(tmpSlice) > 0 {
				res = append(res, tmpSlice)
			}

			newSize := len(q)    //新的q长度
			copy(q[:], q[size:]) //移动q，新追加的内容覆盖以前的内容
			q = q[:newSize-size] //重新规划q的内容
		}
		depth++

		//设置结束条件，到最后一层没有内容
		if len(q) == 0 {
			break
		}
	}
	return
}
