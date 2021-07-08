package answers

//https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
//11.旋转数组的最小数字
func MinArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		mid := left + (right-left)>>1

		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] <= numbers[right] {
			right = right - 1
		}
	}
	return numbers[left]
}
