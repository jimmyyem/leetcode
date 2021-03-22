package answers

//https://leetcode-cn.com/problems/container-with-most-water/
//11. 盛最多水的容器
func MaxArea(height []int) int {
	var max, start, end, tempArea, high, width int
	end = len(height) - 1

	//挨个遍历，获取每个容器的面积大小，最后确定最大面积
	for start < end {
		//确定宽度
		width = end - start

		//根据对比确定高度（木桶原理）
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		//计算面积并与max比较
		tempArea = width * high
		if tempArea > max {
			max = tempArea
		}
	}

	return max
}
