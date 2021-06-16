package answers

//https://leetcode-cn.com/problems/maximum-subarray/
//53.最大子序列
func MaxSubArray(nums []int) int {
	//res, sum := nums[0], 0
	//for i := 0; i < len(nums); i++ {
	//	sum = nums[i] + max(0, sum)
	//	res = max(res, sum)
	//}
	//
	//return res

	/**
	dp[i] 最大之序和


	*/
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
