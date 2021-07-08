package answers

//https://leetcode-cn.com/problems/continuous-subarray-sum/
//523. 连续的子数组和
func CheckSubarraySum(nums []int, k int) bool {
	//sum function
	sumFn := func(num []int) int {
		retVal := 0
		for _, val := range num {
			retVal += val
		}
		return retVal
	}

	for i := 0; i < len(nums)-1; i++ {
		if len(nums[:i]) >= 2 {
			sumVal := sumFn(nums[:i])
			if sumVal%k == 0 {
				return true
			}
		}
	}
	return false
}
