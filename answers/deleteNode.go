package answers

//https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
//237. 删除链表中的节点
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func DeleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
