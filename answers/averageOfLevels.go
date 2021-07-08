package answers

//https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/
//637. 二叉树的层平均值
func AverageOfLevels(root *TreeNode) (res []float64) {
	if root == nil {
		return
	}

	avarage := func(set []int) float64 {
		if len(set) == 0 {
			return 0.0
		}

		sum := 0
		for _, val := range set {
			sum += val
		}

		return float64(sum)/float64(len(set))
	}

	tmpQ := []*TreeNode{root}
	for {
		set := []int{}
		for i := 0; i < len(tmpQ); i++ {
			set = append(set, tmpQ[i].Val)
		}
		res = append(res, avarage(set))

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
