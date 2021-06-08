package main

import (
	"fmt"
)

func main() {
	//字符串类型一维数组排序
	//strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	//sort.Slice(strs, func(i, j int) bool {
	//	return strs[i] < strs[j]
	//})
	//fmt.Printf("%v\n", strs)

	//数字类型一维数组排序
	nums := []int{1, 5, 9, 3, 4, 2, 6, 8, 22}
	//nums = bubbleSort(nums)
	//nums = selectSort(nums)
	//quickSort(nums, 0, len(nums)-1)
	heapSort(nums)
	fmt.Printf("%v\n", nums)

	//sort.Slice(nums, func(i, j int) bool {
	//	return nums[i] < nums[j]
	//})
	//fmt.Printf("%v\n", nums)
}

//冒泡排序
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}

//选择排序
func selectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			//找到最小的和i交换位置
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}

	return nums
}

//插入排序
func insertSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}

	//升序排列
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

	return arr
}

//快速排序
func quickSort(sortArray []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		midValue := sortArray[mid]
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
		quickSort(sortArray, left, i-1)
		quickSort(sortArray, j+1, right)
	}
}

//归并排序
func mergeSort(s []int) []int {
	len := len(s)
	if len == 1 {
		return s //最后切割只剩下一个元素
	}
	m := len / 2
	leftS := mergeSort(s[:m])
	rightS := mergeSort(s[m:])
	return merge(leftS, rightS)
}

//把两个有序切片合并成一个有序切片
func merge(l []int, r []int) []int {
	lLen := len(l)
	rLen := len(r)
	res := make([]int, 0)

	lIndex, rIndex := 0, 0 //两个切片的下标，插入一个数据，下标加一
	for lIndex < lLen && rIndex < rLen {
		if l[lIndex] > r[rIndex] {
			res = append(res, r[rIndex])
			rIndex++
		} else {
			res = append(res, l[lIndex])
			lIndex++
		}
	}
	if lIndex < lLen { //左边的还有剩余元素
		res = append(res, l[lIndex:]...)
	}
	if rIndex < rLen {
		res = append(res, r[rIndex:]...)
	}

	return res
}

//堆排序
//升序
func heapSort(nums []int) {
	fmt.Printf("初始状态 %v\n", nums)
	n := len(nums) - 1

	//从底层向上层构建，n/2是第一个非叶子，叶子节点不需要动，这样就减少了将近一半节点的比较和移动
	//这时候堆顶已经是最大/小的了
	for k := n / 2; k >= 1; k-- {
		sink(nums, k, n)
	}
	fmt.Printf("建堆后 %v\n", nums)

	//首尾交换，重新构建小顶堆
	for n > 1 {
		//将最大/小的数值调整到堆的末尾
		swap(nums, 1, n)
		//减少堆的长度
		n--
		//由于堆顶元素改变了，而且堆的大小改变了，需要重新调整堆，维持堆的性质
		sink(nums, 1, n)
	}
	fmt.Printf("排序后 %v", nums)
}

//下沉操作
func sink(nums []int, k, n int) {
	for {
		i := k * 2	//找到左子节点
		if i > n {
			break
		}
		//如果右节点比左节点大，则使用右节点
		if i < n && nums[i+1] > nums[i] {
			i++
		}
		//如果子节点比父节点小，这交换；否则就停止下沉
		if nums[k] >= nums[i] {
			break
		} else {
			swap(nums, k, i)
		}

		k = i
	}
}

//交换slice的2个元素的值
func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}