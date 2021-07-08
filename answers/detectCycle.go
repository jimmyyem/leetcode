package answers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	if fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return slow
		}
	}

	return nil
}

//合并两个有序链表，并保持有序
func MergeList(headA, headB *ListNode) *ListNode {
	tmpA, tmpB := headA, headB
	tmp := &ListNode{
		Val:  0,
		Next: nil,
	}
	res := tmp

	for tmpA != nil && tmpB != nil {
		if tmpA.Val > tmpB.Val {
			tmp.Next = tmpB
			tmpB = tmpB.Next
		} else {
			tmp.Next = tmpA
			tmpA = tmpA.Next
		}
		tmp = tmp.Next
	}

	if tmpA == nil {
		tmp.Next = tmpB
	} else {
		tmp.Next = tmpA
	}

	return res.Next
}

//删除有序链表中重复节点
func RemoveDupulicateNode(head *ListNode) *ListNode {
	if head == nil && head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	hasSame := false

	for {
		if slow == nil || fast == nil {
			break
		}

		//利用fast发现重复节点
		if fast.Next != nil && fast.Val == fast.Next.Val {
			hasSame = true
			fast = fast.Next
			continue
		}

		//过滤掉重复节点
		if hasSame && fast.Next != nil && fast.Val != fast.Next.Val {
			hasSame = false
			slow.Next = fast.Next	//1,2,3,3,3,4 会剩余 1,2,4
			//slow.Next = fast		//1,2,3,3,3,4 会剩余 1,2,3,4
			fast = fast.Next
			continue
		}

		fast = fast.Next
		slow = slow.Next
	}

	return head
}
