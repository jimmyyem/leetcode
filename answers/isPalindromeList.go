package answers

import "fmt"

//https://leetcode-cn.com/problems/palindrome-linked-list/
//234. 回文链表
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func IsPalindromeList(head *ListNode) bool {
	set := make([]int, 0)
	for head != nil {
		set = append(set, head.Val)
		head = head.Next
	}
	fmt.Printf("%v\n", set)

	left, right := 0, len(set)-1
	for left < right {
		if set[left] != set[right] {
			return false
		}
		left++
		right--
	}

	return true
}
