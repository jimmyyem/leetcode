package main

import (
	"fmt"
	"sort"
)

func main() {
	//字符串类型一维数组排序
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})
	fmt.Printf("%v\n", strs)

	//数字类型一维数组排序
	nums := []int{1, 5, 9, 3, 4, 2, 6, 8, 22}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Printf("%v\n", nums)

	//二维数组排序
	arr := make([][]int, 4)
	arr[0] = []int{5, 10}
	arr[1] = []int{2, 5}
	arr[2] = []int{4, 7}
	arr[3] = []int{3, 9}
	fmt.Printf("%v\n", arr)

	fmt.Println(maximumUnits(arr, 10))

}

//对二维数组排序
//https://leetcode-cn.com/problems/maximum-units-on-a-truck/
//1710. 卡车上的最大单元数
func maximumUnits(boxTypes [][]int, truckSize int) int {
	sum := 0

	//这里是按numberOfUnitsPerBox降序，优先使用大的装载单元数量
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	//遍历boxTypes，直到达到truckSize停止
	for i := 0; i < len(boxTypes); i++ {
		if boxTypes[i][0] < truckSize {
			truckSize -= boxTypes[i][0]
			sum += boxTypes[i][0] * boxTypes[i][1]
		} else {
			sum += truckSize * boxTypes[i][1]
			break
		}
	}

	return sum
}
