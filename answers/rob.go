package answers

//https://leetcode-cn.com/problems/house-robber/
//198. 打家劫舍
func Rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	//dp[i] 代表到达i时能偷窃到的最多金额
	dp := make([]int, len(nums)+1)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}
