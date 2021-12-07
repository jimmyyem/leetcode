package answers

//https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
//剑指 Offer 22. 链表中倒数第k个节点
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func GetKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head

	// 2个指针，fast先走，领先slow指针k-1步，这样fast走到头的时候slow正好位于距离fast距离为k的位置
	for k > 0 && fast != nil {
		k--
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow

	//left := head
	//right := &ListNode{}
	//
	//count := 0
	//tmpList := head
	//for tmpList != nil {
	//	if count == k-1 {
	//		right = tmpList
	//		break
	//	}
	//	count++
	//	tmpList = tmpList.Next
	//}
	//
	//for {
	//	if right.Next == nil {
	//		break
	//	}
	//	right = right.Next
	//	left = left.Next
	//}
	//
	//return left
}
