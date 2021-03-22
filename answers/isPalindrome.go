package answers

import (
	"math"
	"strconv"
)

//https://leetcode-cn.com/problems/palindrome-number/
//9.回文数

func isPalindrome(x int) bool {
	//负数肯定不是回文
	if x < 0 {
		return false
	}
	//0~9肯定是回文
	if x >= 0 && x <= 9 {
		return true
	}

	//转成字符串来处理
	str := strconv.Itoa(x)
	n := len(str)
	for i := 0; i < len(str); i++ {
		if str[i] != str[n-i] {
			return false
		}
	}

	return true
}

func IsPalindrome(x int) bool {
	//负数肯定不是回文
	if x < 0 {
		return false
	}
	//0~9肯定是回文
	if x >= 0 && x <= 9 {
		return true
	}

	var round, temp, suffix, prefix = 0, x, 0, 0
	for temp > 0 {
		round++
		temp /= 10
		if temp == 0 {
			break
		}
	}

	for x > 0 {
		suffix = x % 10
		prefix = x / int(math.Pow10(round-1))

		if suffix != prefix {
			return false
		}

		x = (x - prefix*int(math.Pow10(round-1))) / 10
		round -= 2
	}

	return true
}
