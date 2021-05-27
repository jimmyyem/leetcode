package answers

//https://leetcode-cn.com/problems/reverse-bits/
//190. 颠倒二进制位
func ReverseBits(num uint32) uint32 {
	var res uint32
	pointer := 31

	for num > 0 {
		// 计算最后一位然后左移相应位数
		res += (num & 1) << pointer
		// 左移位数减1
		pointer--
		// 当前数字右移一位
		num = num >> 1
	}

	return res
}
