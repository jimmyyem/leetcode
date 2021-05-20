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

func QuickSort(sortArray []int, left, right int) {
	if left < right {
		midValue := sortArray[(left+right)/2]
		i := left
		j := right

		for {
			for sortArray[i] < midValue {
				i++
			}
			for sortArray[j] > midValue {
				j--
			}
			if i >= j {
				break
			}
			sortArray[i], sortArray[j] = sortArray[j], sortArray[i]
		}
		//fmt.Printf("%v, %d\n", sortArray, i)
		QuickSort(sortArray, left, i-1)
		QuickSort(sortArray, j+1, right)
	}
}

//
func findIndex(nums []int) int {
	var idx int
	if len(nums) > 0 {

	}

	return idx
}
