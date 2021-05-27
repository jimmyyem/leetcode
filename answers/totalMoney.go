package answers

//https://leetcode-cn.com/problems/calculate-money-in-leetcode-bank/
//1716. 计算力扣银行的钱
func TotalMoney(n int) int {
	total := 0
	weeks := int(n / 7)

	getSum := func(start, end int) int {
		return (start + end) * (end - start + 1) / 2
	}

	//不足1周
	if weeks == 0 && n > 0 {
		return getSum(1, n)
	}

	//够一周
	start := 1
	for i := 0; i < weeks; i++ {
		total += getSum(start, start+6)
		start++
	}

	//剩余
	if n%7 != 0 {
		total += getSum(start, start+n%7-1)
	}

	return total
}
