package answers

import (
	"sort"
)

//https://leetcode-cn.com/problems/maximum-units-on-a-truck/
//1710. 卡车上的最大单元数
func MaximumUnits(boxTypes [][]int, truckSize int) int {
	var total, multiplier int
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	for i := 0; i < len(boxTypes); i++ {
		if truckSize > 0 {
			if boxTypes[i][0] <= truckSize {
				multiplier = boxTypes[i][0]
			} else {
				multiplier = truckSize
			}
			total += multiplier * boxTypes[i][1]
			truckSize -= boxTypes[i][0]
		}
	}

	return total
}
