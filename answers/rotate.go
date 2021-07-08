package answers

//https://leetcode-cn.com/problems/rotate-image/
//48. 旋转图像
func Rotate(matrix [][]int) {
	n := len(matrix)

	/**
	旋转操作的结果与先水平翻，后对角线翻结果相同
	*/

	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
func Rotate2(matrix [][]int) {
	size := len(matrix)
	if size == 1 {
		return
	}

	temp := make([][]int, size)
	for k := 0; k < size; k++ {
		temp[k] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			//fmt.Println(i,j, matrix[size-j-1][i])
			temp[i][j] = matrix[size-j-1][i]
		}
	}
	//fmt.Printf("%v\n", temp)
	copy(matrix, temp)
}

/**
0,0 	=> 2,0
1,0 	=> 2,1
2,0 	=> 2,2






*/
