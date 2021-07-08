package answers

//https://leetcode-cn.com/problems/cousins-in-binary-tree/
//993. 二叉树的堂兄弟节点

var (
	depth  = map[int]int{} //节点所在的深度
	parent = map[int]int{} //节点的父节点的值
)

func IsCousins(root *TreeNode, x int, y int) bool {
	var xDepth, yDepth, depth int
	var xParent, yParent int

	nodes := []*TreeNode{root}

	for len(nodes) > 0 {
		tmpNodes := make([]*TreeNode, 0)

		for i := 0; i < len(nodes); i++ {
			//下一级节点加入tmpNode，用于下一次遍历
			if nodes[i].Left != nil {
				tmpNodes = append(tmpNodes, nodes[i].Left)
				if nodes[i].Left.Val == x {
					xParent = nodes[i].Val
				}
				if nodes[i].Left.Val == y {
					yParent = nodes[i].Val
				}
			}
			if nodes[i].Right != nil {
				tmpNodes = append(tmpNodes, nodes[i].Right)
				if nodes[i].Right.Val == x {
					xParent = nodes[i].Val
				}
				if nodes[i].Right.Val == y {
					yParent = nodes[i].Val
				}
			}
			if nodes[i].Val == x && xDepth == 0 {
				xDepth = depth
			}
			if nodes[i].Val == y && yDepth == 0 {
				yDepth = depth
			}
		}

		//如果已经找到，提前结束
		if xDepth > 0 && yDepth > 0 && depth >= 2 {
			break
		}

		//继续下次遍历
		depth++
		nodes = tmpNodes
	}

	return xDepth == yDepth && xParent != yParent

	//dfs(root, 0, 0)
	//return depth[x] == depth[y] && parent[x] != parent[y]
}

func dfs(root *TreeNode, par, deep int) {
	if root == nil {
		return
	}

	depth[root.Val] = deep
	parent[root.Val] = par

	if root.Left != nil {
		dfs(root.Left, root.Val, deep+1)
	}
	if root.Right != nil {
		dfs(root.Right, root.Val, deep+1)
	}
}
