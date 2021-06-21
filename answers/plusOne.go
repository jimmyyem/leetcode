package answers

//https://leetcode-cn.com/problems/plus-one/solution/
//66. 加一
func PlusOne(digits []int) []int {
	flag := false //标识是否进位了
	for i := len(digits) - 1; i >= 0; i-- {
		if flag == false {
			if digits[i] == 9 {
				digits[i] = 0
				flag = true
			} else {
				digits[i]++
				break
			}
		} else {
			if digits[i] < 9 {
				digits[i]++
				flag = false
				break
			} else {
				digits[i] = 0	//这里可能漏掉的情况下面会补齐
			}
		}
	}

	//[9] 这样的情况还要最前面补1
	if flag == true {
		digits = append([]int{1}, digits...)
	}

	return digits
}
