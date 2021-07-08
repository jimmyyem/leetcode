package answers

//https://leetcode-cn.com/problems/find-the-highest-altitude/
//1732. 找到最高海拔
func LargestAltitude(gain []int) int {
	var res, start, next, temp int

	var maxNumber = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(gain); i++ {
		next = gain[i] + start
		temp = maxNumber(start, next)
		if temp > res {
			res = temp
		}
		start = next
	}

	return res
}
