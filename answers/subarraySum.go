package answers

//https://leetcode-cn.com/problems/subarray-sum-equals-k/
//560. 和为K的子数组
func SubarraySum(nums []int, k int) int {
	/**
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
		sum = 0
	}

	return count
	*/

	/**
	详细解释见@see
	@see https://github.com/chefyuan/algorithm-base/blob/main/animation-simulation/%E6%95%B0%E7%BB%84%E7%AF%87/leetcode560%E5%92%8C%E4%B8%BAK%E7%9A%84%E5%AD%90%E6%95%B0%E7%BB%84.md
	原始数组  [1, 2, 6, 4, -1, 5]
	index     	0   1   2   3   4   5
	nums  		1   2   6   4   -1  5
	preSum   	0   1   3   9   13  12  17

	//得出结论
	preSum[i] = preSum[i-1] + nums[i]
	*/

	count := 0
	preSum := make([]int, len(nums)+1) //preSum[0]就是默认值0
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = nums[i] + preSum[i]
	}
	//fmt.Printf("%v\n", preSum)
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if preSum[j+1]-preSum[i] == k {
				count++
			}
		}
	}

	return count
}
