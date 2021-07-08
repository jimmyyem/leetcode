package answers

// 5张扑克牌判断是否连续
// https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
func IsStraight(nums []int) bool {
	mp := make(map[int]int)
	var (
		maxVal int = 0
		minVal int = 14
	)

	for _, value := range nums {
		//fmt.Println(value)
		if value == 0 {
			continue
		}
		//不能重复
		_, found := mp[value]
		if found {
			return false
		}
		mp[value] = 1

		maxVal = max(value, maxVal)
		minVal = min(value, minVal)
	}
	//最大与最小值的差肯定是小于5
	return maxVal-minVal < 5
}
