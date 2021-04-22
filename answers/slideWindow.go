package answers

//
//滑动窗口
func SlidingWindow(size int, input []int) [][]int {
	// 返回入参的切片作为第一个元素
	if len(input) <= size {
		return [][]int{input}
	}

	// 以所需的精确大小分配切片
	r := make([][]int, 0, len(input)-size+1)

	for i, j := 0, size; j <= len(input); i, j = i+1, j+1 {
		r = append(r, input[i:j])
	}

	return r
}
