package answers

import "math"

//https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/
//515. 在每个树行中找最大值
func LargestValues(root *TreeNode) (res []int) {
	if root == nil {
		return
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	tmpQ := []*TreeNode{root}
	for {
		maxV := math.MinInt32
		for i := 0; i < len(tmpQ); i++ {
			maxV = max(maxV, tmpQ[i].Val)
		}
		res = append(res, maxV)

		q := tmpQ
		tmpQ = nil
		for i := 0; i < len(q); i++ {
			if q[i].Left != nil {
				tmpQ = append(tmpQ, q[i].Left)
			}
			if q[i].Right != nil {
				tmpQ = append(tmpQ, q[i].Right)
			}
		}

		if len(tmpQ) == 0 {
			break
		}
	}

	return
}
