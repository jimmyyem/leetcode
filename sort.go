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
	//nums = sortBubble(nums)
	//nums = sortSelect(nums)
	sortQuick(nums, 0, len(nums)-1)
	fmt.Printf("%v\n", nums)

	//sort.Slice(nums, func(i, j int) bool {
	//	return nums[i] < nums[j]
	//})
	//fmt.Printf("%v\n", nums)
}

//冒泡排序
func sortBubble(nums []int) []int {
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
func sortSelect(nums []int) []int {
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
func sortInsert(arr []int) []int {
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
func sortQuick(sortArray []int, left, right int) {
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
		sortQuick(sortArray, left, i-1)
		sortQuick(sortArray, j+1, right)
	}
}
