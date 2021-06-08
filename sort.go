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
	HeapSort(nums)
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
//s[0]不用，实际元素从角标1开始
//父节点元素大于子节点元素
//左子节点角标为2*k
//右子节点角标为2*k+1
//父节点角标为k/2
func HeapSort(s []int) {
	N := len(s) - 1 //s[0]不用，实际元素数量和最后一个元素的角标都为N
	//构造堆
	//如果给两个已构造好的堆添加一个共同父节点，
	//将新添加的节点作一次下沉将构造一个新堆，
	//由于叶子节点都可看作一个构造好的堆，所以
	//可以从最后一个非叶子节点开始下沉，直至
	//根节点，最后一个非叶子节点是最后一个叶子
	//节点的父节点，角标为N/2
	//fmt.Printf("%v\n", s)
	for k := N / 2; k >= 1; k-- {
		Sink(s, k, N)
	}
	//fmt.Printf("%v\n", s)
	//下沉排序
	for N > 1 {
		Swap(s, 1, N) //将大的放在数组后面，升序排序
		N--
		Sink(s, 1, N)
	}
}

//下沉（由上至下的堆有序化）
func Sink(s []int, k, N int) {
	for {
		i := 2 * k //左子节点
		if i > N { //保证该节点是非叶子节点
			break
		}
		if i < N && s[i+1] > s[i] { //选择较大的子节点
			i++
		}
		if s[k] >= s[i] { //没下沉到底就构造好堆了
			break
		} else {
			Swap(s, k, i)
		}

		k = i
	}
}

func Swap(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
}
