package answers

import (
	"fmt"
	"math"
)

//func main() {
//	fmt.Println("5张扑克牌是否连续")
//	nums := []int{0, 0, 0, 19, 5}
//	flags := isStraight(nums)
//	fmt.Println(flags)
//}

// 5张扑克牌判断是否连续
// https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
func IsStraight(nums []int) bool {
	mp := make(map[int]int)
	var max float64 = 0
	var min float64 = 14

	for _, value := range nums {
		fmt.Println(value)
		if value == 0 {
			continue
		}
		_, found := mp[value]
		if found {
			return false
		}
		mp[value] = 1

		max = math.Max(float64(value), max)
		min = math.Min(float64(value), min)
	}
	fmt.Println(mp, max, min)
	return max-min < 5
}
