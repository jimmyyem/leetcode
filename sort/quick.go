package sort

// 快速排序
func Quick(sortArray []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		midValue := sortArray[mid]
		i := left
		j := right

		for {
			// 找到 midValue 左边比 mid 大的值
			for sortArray[i] < midValue {
				i++
			}
			// 找到 midValue 右边比 mid 小的值
			for sortArray[j] > midValue {
				j--
			}
			if i >= j {
				break
			}
			sortArray[i], sortArray[j] = sortArray[j], sortArray[i]
		}
		//fmt.Printf("%v, %d\n", sortArray, i)
		Quick(sortArray, left, i-1)
		Quick(sortArray, j+1, right)
	}
}
