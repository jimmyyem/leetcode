package answers

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode-cn.com/problems/merge-two-sorted-lists/
//21. 合并两个有序链表
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 如果某一个是空的 直接返回另一个
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	res := head //head用于遍历；res用于最后返回的时候用，一直指向表头
	// 如果2个链表有内容就一直遍历
	for list1 != nil && list2 != nil {
		//fmt.Println(list1.Val, list2.Val)
		if list1.Val > list2.Val {
			head.Next = list2
			list2 = list2.Next
		} else {
			head.Next = list1
			list1 = list1.Next
		}
		head = head.Next
	}
	// 直到至少1个链表无内容了 就直接用head接上另一个链表即可
	if list1 != nil {
		head.Next = list1
	}
	if list2 != nil {
		head.Next = list2
	}

	return res.Next
}
