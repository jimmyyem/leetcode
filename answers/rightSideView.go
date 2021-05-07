package answers

//https://leetcode-cn.com/problems/binary-tree-right-side-view/
//199. 二叉树的右视图
func RightSideView(root *TreeNode) (res []int) {
	if root == nil {
		return
	}

	tmpQ := []*TreeNode{root}
	for {
		last := tmpQ[len(tmpQ)-1]
		res = append(res, last.Val)

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
		q = nil

		if len(tmpQ) == 0 {
			break
		}
	}

	return
}
