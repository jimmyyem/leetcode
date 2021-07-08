package answers

//https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
//107. 二叉树的层序遍历
func LevelOrderBottom(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	var depth int
	q := make(map[int][]*TreeNode)
	temp1 := make([]*TreeNode, 0) //用于存放下一层节点
	temp2 := make([]*TreeNode, 0)

	//遍历树，把树的节点挨个遍历并且存储在一个map[int][]*TreeNode里，然后倒序遍历获得
	for {
		if depth == 0 {
			q[depth] = append(q[depth], root)
			if root.Left != nil {
				temp1 = append(temp1, root.Left)
			}
			if root.Right != nil {
				temp1 = append(temp1, root.Right)
			}
		} else {
			for _, val := range temp1 {
				q[depth] = append(q[depth], val)

				if val.Left != nil {
					temp2 = append(temp2, val.Left)
				}
				if val.Right != nil {
					temp2 = append(temp2, val.Right)
				}
			}

			//结束条件
			if len(temp2) == 0 {
				break
			}

			temp1 = append(temp1, temp2...) //给temp扩容
			temp1 = temp1[:len(temp2)]
			copy(temp1, temp2) //从temp2拷贝内容到temp
			temp2 = temp2[:0]  //让temp2内容清空重置
		}

		depth++
	}

	for depth >= 0 {
		if val, found := q[depth]; found {
			tempSlice := make([]int, 0)
			for _, v := range val {
				tempSlice = append(tempSlice, v.Val)
			}
			res = append(res, tempSlice)
			tempSlice = tempSlice[:0]
		}
		depth--
	}

	return
}
