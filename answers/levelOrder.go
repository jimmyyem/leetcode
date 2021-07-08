package answers

//https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
//102. 二叉树的层序遍历
func LevelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		tmpSlice := []int{}
		tmpQ := []*TreeNode{}

		for i := 0; i < len(q); i++ {
			if q[i] != nil {
				tmpSlice = append(tmpSlice, q[i].Val)
				tmpQ = append(tmpQ, q[i].Left, q[i].Right)
			}
		}

		if len(tmpQ) > 0 { //本层有内容才放入res
			res = append(res, tmpSlice)
			q = tmpQ
		} else { //结束条件，这里如果本层没有节点则说明已经结束
			break
		}
	}
	return
}
