package answers

//https://leetcode-cn.com/problems/maximum-subarray/
//53.最大子序列
func MaxSubArray(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	res, sum := nums[0], 0
	for i := 0; i < len(nums); i++ {
		sum = nums[i] + max(0, sum)
		res = max(res, sum)
	}

	return res
}
