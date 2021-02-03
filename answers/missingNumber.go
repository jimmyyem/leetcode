package answers

//https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof
func MissingNumber(nums []int) int {
	res := 0

	// 找出差值不是1的元素
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] != 1 {
			res = nums[i] + 1
			break
		}
	}

	// 如果没找到，则根据首元素判断取0还是n
	if res == 0 {
		if nums[0] == 0 {
			res = nums[len(nums)-1] + 1
		} else {
			res = nums[0] - 1
		}
	}

	return res
}
