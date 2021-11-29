package answers

// https://leetcode-cn.com/problems/middle-of-the-linked-list/
// 876. 链表的中间结点
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	// 快慢指针，快指针一次走2步，慢指针一次走1步，最后快指针到头时慢指针走了一半路程
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
