package answers

//https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
//剑指 Offer 15. 二进制中1的个数
func HammingWeight(num uint32) int {
	var count int = 0

	for num > 0 {
		count += int(num & 1)
		num /= 2
	}

	return count
}
