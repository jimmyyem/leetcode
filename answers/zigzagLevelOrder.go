package answers

//https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
//103. 二叉树的锯齿形层序遍历
func ZigzagLevelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	var depth int
	q := []*TreeNode{root}

	for len(q) > 0 {
		tmpQ := q
		q = nil

		tmpSlice := make([]int, 0)

		//正常把内容入队列
		for i := 0; i < len(tmpQ); i++ {
			if tmpQ[i].Left != nil {
				q = append(q, tmpQ[i].Left)
			}
			if tmpQ[i].Right != nil {
				q = append(q, tmpQ[i].Right)
			}
		}

		//偶数正常取出，正常入队列q
		if depth%2 == 0 {
			for i := 0; i < len(tmpQ); i++ {
				tmpSlice = append(tmpSlice, tmpQ[i].Val)
			}
		} else { //奇数需要倒叙取出内容
			for i := len(tmpQ) - 1; i >= 0; i-- {
				tmpSlice = append(tmpSlice, tmpQ[i].Val)
			}
		}

		if len(tmpSlice) > 0 { //本层有内容才放入res
			res = append(res, tmpSlice)
		}
		if len(q) == 0 {
			break
		}
		depth++
	}

	return
}
