package answers

//https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
//19. 删除链表的倒数第 N 个结点
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	size, count := 0, 0
	prev := &ListNode{}

	temp := head
	for temp != nil {
		size++
		temp = temp.Next
	}
	//如果只有一个元素肯定是删它，所以返回空
	if size == 1 && n == 1 {
		return nil
	}
	if n > size {
		return head
	}
	if n == size {
		return head.Next
	}

	temp = head
	for temp != nil {
		if count == size-n {
			prev.Next = temp.Next
		} else {
			prev = temp
		}

		count++
		temp = temp.Next
	}

	return head
}
