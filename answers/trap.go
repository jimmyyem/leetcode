package answers

//https://leetcode-cn.com/problems/trapping-rain-water/
//42.接雨水
func Trap(height []int) int {
	i, j := 0, len(height)-1
	leftMax, rightMax := 0, 0
	ans := 0
	for i < j {
		// 移动任何一侧指针，首先计算出指针遍历元素中最大的元素值
		leftMax = max(leftMax, height[i])
		rightMax = max(rightMax, height[j])

		// 每移动一格指针便统计盛水高度
		if height[i] < height[j] {
			ans += leftMax - height[i]
			i++
		} else {
			ans += rightMax - height[j]
			j--
		}
	}
	return ans
}
