package answers

func IsHappy(n int) bool {
	var (
		slow int = n
		fast int = getNext(n)
	)

	//我们不是只跟踪链表中的一个值，而是跟踪两个值，称为快跑者和慢跑者。
	//在算法的每一步中，慢速在链表中前进 1 个节点，快跑者前进 2 个节点（对 getNext(n) 函数的嵌套调用）。
	//如果 n 是一个快乐数，即没有循环，那么快跑者最终会比慢跑者先到达数字 1。
	//如果 n 不是一个快乐的数字，那么最终快跑者和慢跑者将在同一个数字上相遇。
	for fast != slow && fast != 1 {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}

	return fast == 1
}

func getNext(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}

	return sum
}
