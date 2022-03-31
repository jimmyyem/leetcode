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

	/**
	dp := make([]int, len(nums)+1)
	dp[0] = nums[0]
	mx := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > mx {
			mx = dp[i]
		}
	}
	fmt.Printf("%v\n", dp)
	return mx


	*/

	// 改进后的dp
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		temp := 0
		if nums[i]+nums[i-1] > nums[i] {
			temp = nums[i] + nums[i-1]
		}
		if temp > max {
			max = temp
		}
	}
	return max
}
