package answers

//https://leetcode-cn.com/problems/reverse-linked-list/
//206. 反转链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	cur := head
	for cur != nil {
		// 1.（保存一下前进方向）保存下一跳
		temp := cur.Next
		// 2.斩断过去,不忘前事
		cur.Next = prev
		// 3.前驱指针的使命在上面已经完成，这里需要更新前驱指针
		prev = cur
		// 当前指针的使命已经完成，需要继续前进了
		cur = temp
	}

	return prev
}
