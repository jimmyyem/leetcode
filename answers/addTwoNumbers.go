package answers

import (
	"math"
)

//https://leetcode-cn.com/problems/add-two-numbers/
//2.两数相加
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	left, right, deep, sum := 0, 0, 0, 0

	for l1 != nil {
		left += l1.Val * int(math.Pow10(deep))
		deep++
		l1 = l1.Next
	}
	deep = 0
	for l2 != nil {
		right += l2.Val * int(math.Pow10(deep))
		deep++
		l2 = l2.Next
	}

	start := &ListNode{
		Val:  0,
		Next: nil,
	}
	head := start
	sum = left + right

	if sum > 0 {
		for sum > 0 {
			start.Next = &ListNode{
				Val:  sum % 10,
				Next: nil,
			}
			start = start.Next
			sum /= 10
		}
	} else {
		head.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
	}

	return head.Next
}
