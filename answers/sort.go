package answers



//https://www.bookstack.cn/read/For-learning-Go-Tutorial/spilt.4.src-chapter16-01.0.md
//插入排序
func InsertSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}

	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
}
