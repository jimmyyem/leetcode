package answers

// https://leetcode-cn.com/problems/reorder-list/
// 143. 重排链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReorderList(head *ListNode) {
	nodeList := make([]*ListNode,0,8)

	for node := head; node != nil; node = node.Next {
		nodeList = append(nodeList, node)
	}

	start, end := 0, len(nodeList)-1
	for start < end {
		// 0=>n, 1=>n-1
		nodeList[start].Next = nodeList[end]
		start++

		// start++ end-- 时候node可能重叠
		if start == end {
			break
		}

		// n=>1, n-1=>2
		nodeList[end].Next = nodeList[start]
		end--
	}
	nodeList[start].Next = nil
	return
}
