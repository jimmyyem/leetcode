package answers

import "sort"

//https://leetcode-cn.com/problems/merge-intervals/
//88.合并区间
func MergeSection(intervals [][]int) [][]int {
	sort.Slice(intervals, func(x, y int) bool {
		return intervals[x][0] < intervals[y][0]
	})
	res := make([][]int, 0)
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] { //无交集
			res = append(res, prev)
			prev = cur
		} else { //有交集
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
