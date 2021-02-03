package answers

//https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
func FindRepeatNumber(nums []int) int {
	var list = make([]int, len(nums))

	for _, val := range nums {
		list[val]++
	}

	for key, val := range list {
		if val > 1 {
			return key
		}
	}

	return 0
}
