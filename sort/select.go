package sort

// 选择排序
func Select(arr []int) {
	for i := 0; i < len(arr); i++ {
		miniIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[miniIdx] > arr[j] {
				miniIdx = j
			}
		}

		arr[i], arr[miniIdx] = arr[miniIdx], arr[i]
	}
}
