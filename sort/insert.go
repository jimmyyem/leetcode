package sort

// 插入排序
func Insert(array []int) {
	n := len(array)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
}
