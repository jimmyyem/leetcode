package answers

import (
	"math"
	"sort"
)

//https://leetcode-cn.com/problems/3sum-closest/
//16.最接近的3数之和
func ThreeSumClosest(nums []int, target int) int {
	var (
		res  = 0
		n    = len(nums)
		diff = math.MaxInt32
	)

	if n > 2 {
		sort.Ints(nums)
		for i := 0; i < n-2; i++ { //不能到最后一个数要留给j和k
			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]

				//diff最初始是个很大的数，第一次比较后就赋值为1st的差值
				//之后的比较只要diff比差值大，就继续赋值 所以留到最后diff就是最小差值
				if diff > abs(sum-target) {
					res, diff = sum, abs(sum-target)
				}
				if sum == target { //精确找到直接返回
					return sum
				} else if sum < target { //结果比较小
					j++
				} else { //结果比较大
					k--
				}
			}
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
