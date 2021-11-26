package answers

import "math"

//https://leetcode-cn.com/problems/string-to-integer-atoi/
//8.字符串转整数
func MyAtoi(str string) int {
	res, flag := 0, 1
	end, firstNumber := false, false

	for _, val := range str {
		switch val {
		case '-':
			if firstNumber {
				end = true
			} else {
				firstNumber = true
				flag = -1
			}
		case '+':
			if firstNumber {
				end = true
			} else {
				firstNumber = true
			}
		case ' ':
			if firstNumber {
				end = true
			}
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			res = res*10 + int(val-'0')
			firstNumber = true

			//每次都要判断，及早发现溢出的情况
			if res*flag > math.MaxInt32 {
				return math.MaxInt32
			}
			if res*flag < math.MinInt32 {
				return math.MinInt32
			}
		default:
			end = true
		}
		if end {
			break
		}
	}

	return res * flag
}
