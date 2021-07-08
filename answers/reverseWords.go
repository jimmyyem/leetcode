package answers

//https://leetcode-cn.com/problems/reverse-words-in-a-string/solution/
//151. 翻转字符串里的单词
func ReverseWords(s string) string {
	//字符串反转
	reverseWord := func(word []byte) {
		left, right := 0, len(word)-1
		for left < right {
			word[left], word[right] = word[right], word[left]
			left++
			right--
		}
	}

	//去除空格函数
	trimBothWork := func(s string) string {
		left, right := 0, len(s)-1
		for left < right {
			if s[left] == ' ' {
				left++
			} else {
				break
			}
		}
		s = s[left:]

		right = len(s) - 1
		for right >= 0 {
			if s[right] == ' ' {
				right--
			} else {
				break
			}
		}
		s = s[:right+1]

		return string(s)
	}

	s = trimBothWork(s)

	var word = make([]byte, 0)
	var res = make([]byte, 0, len(s))
	for idx := len(s) - 1; idx >= 0; idx-- {
		if s[idx] != ' ' { //非空格，则把内容存起来
			word = append(word, s[idx])
		} else { //空格，把存起来的单词反转并存入res
			reverseWord(word) //把单词逆序
			//必须有内容才增加空格，应对处理多个连续空格的情况
			if len(word) > 0 {
				res = append(res, word...)
				res = append(res, ' ')
			}
			word = word[0:0] //word字节序列重置
		}
	}
	//最后一个单词在这里加上
	if len(word) > 0 {
		reverseWord(word)
		res = append(res, word...)
	}

	return string(res)
}
