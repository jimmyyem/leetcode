package answers

//https://leetcode-cn.com/problems/plus-one/solution/
//66. 加一
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits2 := make([]int, len(digits)+1)
	digits2[0] = 1
	copy(digits2[1:], digits)

	return digits2
}
