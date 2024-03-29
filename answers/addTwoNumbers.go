package answers

//https://leetcode-cn.com/problems/add-two-numbers/
//2.两数相加
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{0, nil}
	//这里用一个result，只是为了后面返回节点方便，并无他用
	result := list

	tmp := 0
	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		// tmp%10 只记录个位数
		list.Next = &ListNode{tmp % 10, nil}
		// 超过10的保留1，下次记录到高位上
		tmp = tmp / 10
		// list列表后移
		list = list.Next
	}

	// 通过result返回，result目前指向列表头
	return result.Next
}
