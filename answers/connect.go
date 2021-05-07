package answers

import "fmt"

//https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/
//116. 填充每个节点的下一个右侧节点指针
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	tmpQ := []*Node{root}

	for len(tmpQ) > 0 {
		q := tmpQ
		tmpQ = nil

		for i := 0; i < len(q); i++ {
			if q[i].Left != nil {
				tmpQ = append(tmpQ, q[i].Left)
			}
			if q[i].Right != nil {
				tmpQ = append(tmpQ, q[i].Right)
			}

			if i+1 < len(q) {
				q[i].Next = q[i+1]
			}
		}
	}

	//tmpQ = []*Node{root}
	//for len(tmpQ) > 0 {
	//	q := tmpQ
	//	tmpQ = nil
	//	for i := 0; i < len(q); i++ {
	//		show(q[i])
	//		if q[i].Left != nil {
	//			tmpQ = append(tmpQ, q[i].Left)
	//		}
	//		if q[i].Right != nil {
	//			tmpQ = append(tmpQ, q[i].Right)
	//		}
	//
	//		if i+1 < len(q) {
	//			q[i].Next = q[i+1]
	//		}
	//	}
	//	fmt.Println(len(tmpQ))
	//}

	return root
}

func show(node *Node) {
	fmt.Printf("Val:%d, Left:%p, Right:%p, Next:%p\n", node.Val, node.Left, node.Right, node.Next)
}
