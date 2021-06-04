package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 5, 6, 8, 12, 14, 16}
	//rest := binarySearch(nums, 8)
	rest := binarySearchPos(nums, 7)
	fmt.Println(rest)
}

//二分查找
//返回值：-1 没查到；其他 >-1 代表查到
func binarySearch(nums []int, k int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == k {
			return mid
		} else if nums[mid] < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func binarySearchPos(nums []int, k int) int {
	left, right := 0, len(nums)-1

	//二分查找
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == k {
			return mid
		} else if nums[mid] < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
